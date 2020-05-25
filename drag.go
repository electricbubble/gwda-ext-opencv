package gwda_ext_opencv

func (sExt *SessionExt) Drag(pathname string, toX, toY int, pressForDuration ...float64) (err error) {
	return sExt.DragFloat(pathname, float64(toX), float64(toY), pressForDuration...)
}

func (sExt *SessionExt) DragFloat(pathname string, toX, toY float64, pressForDuration ...float64) (err error) {
	return sExt.DragOffsetFloat(pathname, toX, toY, 0.5, 0.5, pressForDuration...)
}

func (sExt *SessionExt) DragOffset(pathname string, toX, toY int, xOffset, yOffset float64, pressForDuration ...float64) (err error) {
	return sExt.DragOffsetFloat(pathname, float64(toX), float64(toY), xOffset, yOffset, pressForDuration...)
}

func (sExt *SessionExt) DragOffsetFloat(pathname string, toX, toY, xOffset, yOffset float64, pressForDuration ...float64) (err error) {
	if len(pressForDuration) == 0 {
		pressForDuration = []float64{1.0}
	}

	var x, y, width, height float64
	if x, y, width, height, err = sExt.FindImageRectInUIKit(pathname); err != nil {
		return err
	}

	fromX := x + width*xOffset
	fromY := y + height*yOffset

	return sExt.s.DragFloat(fromX, fromY, toX, toY, pressForDuration[0])
}
