package gwda_ext_opencv

import (
	"github.com/electricbubble/gwda"
)

func (dExt *DriverExt) Tap(pathname string) error {
	return dExt.TapOffset(pathname, 0.5, 0.5)
}

func (dExt *DriverExt) TapOffset(pathname string, xOffset, yOffset float64) (err error) {
	var x, y, width, height float64
	if x, y, width, height, err = dExt.FindImageRectInUIKit(pathname); err != nil {
		return err
	}

	return dExt.driver.TapFloat(x+width*xOffset, y+height*yOffset)
}

func (dExt *DriverExt) DoubleTap(pathname string) (err error) {
	return dExt.DoubleTapOffset(pathname, 0.5, 0.5)
}

func (dExt *DriverExt) DoubleTapOffset(pathname string, xOffset, yOffset float64) (err error) {
	var x, y, width, height float64
	if x, y, width, height, err = dExt.FindImageRectInUIKit(pathname); err != nil {
		return err
	}

	return dExt.driver.DoubleTapFloat(x+width*xOffset, y+height*yOffset)
}

// TapWithNumber sends one or more taps
func (dExt *DriverExt) TapWithNumber(pathname string, numberOfTaps int) (err error) {
	return dExt.TapWithNumberOffset(pathname, numberOfTaps, 0.5, 0.5)
}

func (dExt *DriverExt) TapWithNumberOffset(pathname string, numberOfTaps int, xOffset, yOffset float64) (err error) {
	if numberOfTaps <= 0 {
		numberOfTaps = 1
	}
	var x, y, width, height float64
	if x, y, width, height, err = dExt.FindImageRectInUIKit(pathname); err != nil {
		return err
	}

	x = x + width*xOffset
	y = y + height*yOffset

	touchActions := gwda.NewTouchActions().Tap(gwda.NewTouchActionTap().WithXYFloat(x, y).WithCount(numberOfTaps))
	return dExt.PerformTouchActions(touchActions)
}
