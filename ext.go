package gwda_ext_opencv

import (
	"github.com/electricbubble/gwda"
	cvHelper "github.com/electricbubble/opencv-helper"
	"image"
	"path/filepath"
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

type SessionExt struct {
	s         *gwda.Session
	pathname  string
	scale     float64
	MatchMode TemplateMatchMode
}

// Extend 获得扩展后的 Session，
// 并指定截图保存的路径，
// 获取当前设备的 Scale，
// 默认匹配模式为 TmCcoeffNormed，
// 默认关闭匹配值的输出
func Extend(session *gwda.Session, pathname string, matchMode ...TemplateMatchMode) (sExt *SessionExt, err error) {
	sExt = &SessionExt{s: session}
	if err = cvHelper.StoreDirectory(pathname); err != nil {
		return nil, err
	}
	sExt.pathname = pathname

	if sExt.scale, err = sExt.s.Scale(); err != nil {
		return &SessionExt{}, err
	}

	if len(matchMode) == 0 {
		matchMode = []TemplateMatchMode{TmCcoeffNormed}
	}
	sExt.MatchMode = matchMode[0]
	cvHelper.Debug(cvHelper.DebugMode(DmOff))
	return sExt, nil
}

func (sExt *SessionExt) Debug(dm DebugMode) {
	cvHelper.Debug(cvHelper.DebugMode(dm))
}

func (sExt *SessionExt) findImgRect(search string, confidence float64) (rect image.Rectangle, err error) {
	pathSource := filepath.Join(sExt.pathname, cvHelper.GenFilename())
	if err = sExt.s.ScreenshotToDisk(pathSource); err != nil {
		return image.Rectangle{}, err
	}

	if rect, err = cvHelper.FindImageRectFromDisk(pathSource, search, float32(confidence), cvHelper.TemplateMatchMode(sExt.MatchMode)); err != nil {
		return image.Rectangle{}, err
	}
	return
}

func (sExt *SessionExt) mappingToRectInUIKit(rect image.Rectangle) (x, y, width, height float64) {
	x, y = float64(rect.Min.X)/sExt.scale, float64(rect.Min.Y)/sExt.scale
	width, height = float64(rect.Dx())/sExt.scale, float64(rect.Dy())/sExt.scale
	return
}

// IsExist
