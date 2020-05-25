package gwda_ext_opencv

import (
	"github.com/electricbubble/gwda"
)

func (sExt *SessionExt) Tap(pathname string) error {
	return sExt.TapOffset(pathname, 0.5, 0.5)
}

func (sExt *SessionExt) TapOffset(pathname string, xOffset, yOffset float64) (err error) {
	var x, y, width, height float64
	if x, y, width, height, err = sExt.FindImageRectInUIKit(pathname); err != nil {
		return err
	}

	return sExt.s.TapFloat(x+width*xOffset, y+height*yOffset)
}

func (sExt *SessionExt) DoubleTap(pathname string) (err error) {
	return sExt.DoubleTapOffset(pathname, 0.5, 0.5)
}

func (sExt *SessionExt) DoubleTapOffset(pathname string, xOffset, yOffset float64) (err error) {
	var x, y, width, height float64
	if x, y, width, height, err = sExt.FindImageRectInUIKit(pathname); err != nil {
		return err
	}

	return sExt.s.DoubleTapFloat(x+width*xOffset, y+height*yOffset)
}

// TapWithNumber sends one or more taps
func (sExt *SessionExt) TapWithNumber(pathname string, numberOfTaps int) (err error) {
	return sExt.TapWithNumberOffset(pathname, numberOfTaps, 0.5, 0.5)
}

func (sExt *SessionExt) TapWithNumberOffset(pathname string, numberOfTaps int, xOffset, yOffset float64) (err error) {
	if numberOfTaps <= 0 {
		numberOfTaps = 1
	}
	var x, y, width, height float64
	if x, y, width, height, err = sExt.FindImageRectInUIKit(pathname); err != nil {
		return err
	}

	x = x + width*xOffset
	y = y + height*yOffset

	touchActions := gwda.NewWDATouchActions().Tap(gwda.NewWDATouchActionOptionTap().SetXYFloat(x, y).SetCount(numberOfTaps))
	return sExt.PerformTouchActions(touchActions)
}
