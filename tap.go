package gwda_ext_opencv

import "image"

func (sExt *SessionExt) Tap(pathname string, confidence float64) error {
	return sExt.TapOffset(pathname, confidence, 0.5, 0.5)
}

func (sExt *SessionExt) TapOffset(pathname string, confidence float64, xOffset, yOffset float64) (err error) {
	var rect image.Rectangle
	if rect, err = sExt.findImgRect(pathname, confidence); err != nil {
		return err
	}

	x, y, width, height := sExt.mappingToRectInUIKit(rect)

	return sExt.s.TapFloat(x+width*xOffset, y+height*yOffset)
}

func (sExt *SessionExt) DoubleTap(pathname string, confidence float64) (err error) {
	return sExt.DoubleTapOffset(pathname, confidence, 0.5, 0.5)
}

func (sExt *SessionExt) DoubleTapOffset(pathname string, confidence float64, xOffset, yOffset float64) (err error) {
	var rect image.Rectangle
	if rect, err = sExt.findImgRect(pathname, confidence); err != nil {
		return err
	}

	x, y, width, height := sExt.mappingToRectInUIKit(rect)

	return sExt.s.DoubleTapFloat(x+width*xOffset, y+height*yOffset)
}

// TODO ForceTouch
// TODO TouchAndHoldFloat
// TODO Drag
// TODO Pinch
// TODO TwoFingerTap
