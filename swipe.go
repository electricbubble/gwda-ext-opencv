package gwda_ext_opencv

func (sExt *SessionExt) SwipeUp(pathname string, distance ...float64) (err error) {
	return sExt.SwipeUpOffset(pathname, 0.5, 0.9, distance...)
}

func (sExt *SessionExt) SwipeUpOffset(pathname string, xOffset, yOffset float64, distance ...float64) (err error) {
	if len(distance) == 0 {
		distance = []float64{1.0}
	}

	var x, y, width, height float64
	if x, y, width, height, err = sExt.FindImgRectInUIKit(pathname); err != nil {
		return err
	}

	fromX := x + width*xOffset
	fromY := (y + height) - height*(1.0-yOffset)

	toX := fromX
	toY := fromY - height*distance[0]

	return sExt.s.SwipeFloat(fromX, fromY, toX, toY)
}

func (sExt *SessionExt) SwipeDown(pathname string, distance ...float64) (err error) {
	return sExt.SwipeDownOffset(pathname, 0.5, 0.1, distance...)
}

func (sExt *SessionExt) SwipeDownOffset(pathname string, xOffset, yOffset float64, distance ...float64) (err error) {
	if len(distance) == 0 {
		distance = []float64{1.0}
	}

	var x, y, width, height float64
	if x, y, width, height, err = sExt.FindImgRectInUIKit(pathname); err != nil {
		return err
	}

	fromX := x + width*xOffset
	fromY := y + height*yOffset

	toX := fromX
	toY := fromY + height*distance[0]

	return sExt.s.SwipeFloat(fromX, fromY, toX, toY)
}

func (sExt *SessionExt) SwipeLeft(pathname string, distance ...float64) (err error) {
	return sExt.SwipeLeftOffset(pathname, 0.9, 0.5, distance...)
}

func (sExt *SessionExt) SwipeLeftOffset(pathname string, xOffset, yOffset float64, distance ...float64) (err error) {
	if len(distance) == 0 {
		distance = []float64{1.0}
	}

	var x, y, width, height float64
	if x, y, width, height, err = sExt.FindImgRectInUIKit(pathname); err != nil {
		return err
	}

	fromX := x + width*xOffset
	fromY := y + height*yOffset

	toX := fromX - width*distance[0]
	toY := fromY

	return sExt.s.SwipeFloat(fromX, fromY, toX, toY)
}

func (sExt *SessionExt) SwipeRight(pathname string, distance ...float64) (err error) {
	return sExt.SwipeRightOffset(pathname, 0.1, 0.5, distance...)
}

func (sExt *SessionExt) SwipeRightOffset(pathname string, xOffset, yOffset float64, distance ...float64) (err error) {
	if len(distance) == 0 {
		distance = []float64{1.0}
	}

	var x, y, width, height float64
	if x, y, width, height, err = sExt.FindImgRectInUIKit(pathname); err != nil {
		return err
	}

	fromX := x + width*xOffset
	fromY := y + height*yOffset

	toX := fromX + width*distance[0]
	toY := fromY

	return sExt.s.SwipeFloat(fromX, fromY, toX, toY)
}
