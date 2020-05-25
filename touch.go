package gwda_ext_opencv

func (sExt *SessionExt) ForceTouch(pathname string, pressure float64, duration ...float64) (err error) {
	return sExt.ForceTouchOffset(pathname, pressure, 0.5, 0.5, duration...)
}

func (sExt *SessionExt) ForceTouchOffset(pathname string, pressure, xOffset, yOffset float64, duration ...float64) (err error) {
	if len(duration) == 0 {
		duration = []float64{1.0}
	}
	var x, y, width, height float64
	if x, y, width, height, err = sExt.FindImageRectInUIKit(pathname); err != nil {
		return err
	}

	return sExt.s.ForceTouchFloat(x+width*xOffset, y+height*yOffset, pressure, duration[0])
}

func (sExt *SessionExt) TouchAndHold(pathname string, duration ...float64) (err error) {
	return sExt.TouchAndHoldOffset(pathname, 0.5, 0.5, duration...)
}

func (sExt *SessionExt) TouchAndHoldOffset(pathname string, xOffset, yOffset float64, duration ...float64) (err error) {
	if len(duration) == 0 {
		duration = []float64{1.0}
	}
	var x, y, width, height float64
	if x, y, width, height, err = sExt.FindImageRectInUIKit(pathname); err != nil {
		return err
	}

	return sExt.s.TouchAndHoldFloat(x+width*xOffset, y+height*yOffset, duration[0])
}
