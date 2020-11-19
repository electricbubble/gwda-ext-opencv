package gwda_ext_opencv

import (
	"bytes"
	"errors"
	"github.com/electricbubble/gwda"
	cvHelper "github.com/electricbubble/opencv-helper"
	"image"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

// TemplateMatchMode is the type of the template matching operation.
type TemplateMatchMode int

const (
	// TmSqdiff maps to TM_SQDIFF
	TmSqdiff TemplateMatchMode = iota
	// TmSqdiffNormed maps to TM_SQDIFF_NORMED
	TmSqdiffNormed
	// TmCcorr maps to TM_CCORR
	TmCcorr
	// TmCcorrNormed maps to TM_CCORR_NORMED
	TmCcorrNormed
	// TmCcoeff maps to TM_CCOEFF
	TmCcoeff
	// TmCcoeffNormed maps to TM_CCOEFF_NORMED
	TmCcoeffNormed
)

type DebugMode int

const (
	// DmOff no output
	DmOff DebugMode = iota
	// DmEachMatch output matched and mismatched values
	DmEachMatch
	// DmNotMatch output only values that do not match
	DmNotMatch
)

type DriverExt struct {
	driver          gwda.WebDriver
	scale           float64
	MatchMode       TemplateMatchMode
	Threshold       float64
	frame           *bytes.Buffer
	doneMjpegStream chan bool
}

// Extend 获得扩展后的 Driver，
// 并指定匹配阀值，
// 获取当前设备的 Scale，
// 默认匹配模式为 TmCcoeffNormed，
// 默认关闭 OpenCV 匹配值计算后的输出
func Extend(driver gwda.WebDriver, threshold float64, matchMode ...TemplateMatchMode) (dExt *DriverExt, err error) {
	dExt = &DriverExt{driver: driver}
	dExt.doneMjpegStream = make(chan bool, 1)

	if dExt.scale, err = dExt.driver.Scale(); err != nil {
		return &DriverExt{}, err
	}

	if len(matchMode) == 0 {
		matchMode = []TemplateMatchMode{TmCcoeffNormed}
	}
	dExt.MatchMode = matchMode[0]
	cvHelper.Debug(cvHelper.DebugMode(DmOff))
	dExt.Threshold = threshold
	return dExt, nil
}

func (dExt *DriverExt) OnlyOnceThreshold(threshold float64) (newExt *DriverExt) {
	newExt = new(DriverExt)
	newExt.driver = dExt.driver
	newExt.scale = dExt.scale
	newExt.MatchMode = dExt.MatchMode
	newExt.Threshold = threshold
	return
}

func (dExt *DriverExt) OnlyOnceMatchMode(matchMode TemplateMatchMode) (newExt *DriverExt) {
	newExt = new(DriverExt)
	newExt.driver = dExt.driver
	newExt.scale = dExt.scale
	newExt.MatchMode = matchMode
	newExt.Threshold = dExt.Threshold
	return
}

func (dExt *DriverExt) Debug(dm DebugMode) {
	cvHelper.Debug(cvHelper.DebugMode(dm))
}

func (dExt *DriverExt) ConnectMjpegStream(httpClient *http.Client) (err error) {
	if httpClient == nil {
		return errors.New(`'httpClient' can't be nil`)
	}

	var req *http.Request
	if req, err = http.NewRequest(http.MethodGet, "http://*", nil); err != nil {
		return err
	}

	var resp *http.Response
	if resp, err = httpClient.Do(req); err != nil {
		return err
	}
	// defer func() { _ = resp.Body.Close() }()

	var boundary string
	if _, param, err := mime.ParseMediaType(resp.Header.Get("Content-Type")); err != nil {
		return err
	} else {
		boundary = strings.Trim(param["boundary"], "-")
	}

	mjpegReader := multipart.NewReader(resp.Body, boundary)

	go func() {
		for {
			select {
			case <-dExt.doneMjpegStream:
				_ = resp.Body.Close()
				return
			default:
				var part *multipart.Part
				if part, err = mjpegReader.NextPart(); err != nil {
					dExt.frame = nil
					continue
				}

				raw := new(bytes.Buffer)
				if _, err = raw.ReadFrom(part); err != nil {
					dExt.frame = nil
					continue
				}
				dExt.frame = raw
			}
		}
	}()

	return
}

func (dExt *DriverExt) CloseMjpegStream() {
	dExt.doneMjpegStream <- true
}

func (dExt *DriverExt) takeScreenshot() (raw *bytes.Buffer, err error) {
	if dExt.frame != nil {
		return dExt.frame, nil
	}
	if raw, err = dExt.driver.Screenshot(); err != nil {
		return nil, err
	}
	return
}

// func (sExt *DriverExt) findImgRect(search string) (rect image.Rectangle, err error) {
// 	pathSource := filepath.Join(sExt.pathname, cvHelper.GenFilename())
// 	if err = sExt.driver.ScreenshotToDisk(pathSource); err != nil {
// 		return image.Rectangle{}, err
// 	}
//
// 	if rect, err = cvHelper.FindImageRectFromDisk(pathSource, search, float32(sExt.Threshold), cvHelper.TemplateMatchMode(sExt.MatchMode)); err != nil {
// 		return image.Rectangle{}, err
// 	}
// 	return
// }

func (dExt *DriverExt) FindAllImageRect(search string) (rects []image.Rectangle, err error) {
	var bufSource, bufSearch *bytes.Buffer
	if bufSearch, err = getBufFromDisk(search); err != nil {
		return nil, err
	}
	if bufSource, err = dExt.takeScreenshot(); err != nil {
		return nil, err
	}

	if rects, err = cvHelper.FindAllImageRectsFromRaw(bufSource, bufSearch, float32(dExt.Threshold), cvHelper.TemplateMatchMode(dExt.MatchMode)); err != nil {
		return nil, err
	}
	return
}

func getBufFromDisk(name string) (*bytes.Buffer, error) {
	var f *os.File
	var err error
	if f, err = os.Open(name); err != nil {
		return nil, err
	}
	var all []byte
	if all, err = ioutil.ReadAll(f); err != nil {
		return nil, err
	}
	return bytes.NewBuffer(all), nil
}

func (dExt *DriverExt) FindImageRectInUIKit(search string) (x, y, width, height float64, err error) {
	var bufSource, bufSearch *bytes.Buffer
	if bufSearch, err = getBufFromDisk(search); err != nil {
		return 0, 0, 0, 0, err
	}
	if bufSource, err = dExt.takeScreenshot(); err != nil {
		return 0, 0, 0, 0, err
	}

	var rect image.Rectangle
	if rect, err = cvHelper.FindImageRectFromRaw(bufSource, bufSearch, float32(dExt.Threshold), cvHelper.TemplateMatchMode(dExt.MatchMode)); err != nil {
		return 0, 0, 0, 0, err
	}

	// if rect, err = dExt.findImgRect(search); err != nil {
	// 	return 0, 0, 0, 0, err
	// }
	x, y, width, height = dExt.MappingToRectInUIKit(rect)
	return
}

func (dExt *DriverExt) MappingToRectInUIKit(rect image.Rectangle) (x, y, width, height float64) {
	x, y = float64(rect.Min.X)/dExt.scale, float64(rect.Min.Y)/dExt.scale
	width, height = float64(rect.Dx())/dExt.scale, float64(rect.Dy())/dExt.scale
	return
}

func (dExt *DriverExt) PerformTouchActions(touchActions *gwda.TouchActions) error {
	return dExt.driver.PerformAppiumTouchActions(touchActions)
}

func (dExt *DriverExt) PerformActions(actions *gwda.W3CActions) error {
	return dExt.driver.PerformW3CActions(actions)
}

// IsExist
