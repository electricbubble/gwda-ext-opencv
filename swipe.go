package gwda_ext_opencv

import "image"

func (sExt *SessionExt) SwipeUp(pathname string, confidence float64, distance ...float64) (err error) {
	return sExt.SwipeUpOffset(pathname, confidence, 0.5, 0.9, distance...)
}

func (sExt *SessionExt) SwipeUpOffset(pathname string, confidence, xOffset, yOffset float64, distance ...float64) (err error) {
	if len(distance) == 0 {
		distance = []float64{1.0}
	}
	var rect image.Rectangle
	if rect, err = sExt.findImgRect(pathname, confidence); err != nil {
		return err
	}

	x, y, width, height := sExt.mappingToRectInUIKit(rect)

	fromX := x + width*xOffset
	fromY := (y + height) - height*(1.0-yOffset)

	toX := fromX
	toY := fromY - height*distance[0]

	return sExt.s.SwipeFloat(fromX, fromY, toX, toY)
}

func (sExt *SessionExt) SwipeDown(pathname string, confidence float64, distance ...float64) (err error) {
	return sExt.SwipeDownOffset(pathname, confidence, 0.5, 0.1, distance...)
}

func (sExt *SessionExt) SwipeDownOffset(pathname string, confidence, xOffset, yOffset float64, distance ...float64) (err error) {
	if len(distance) == 0 {
		distance = []float64{1.0}
	}
	var rect image.Rectangle
	if rect, err = sExt.findImgRect(pathname, confidence); err != nil {
		return err
	}

	x, y, width, height := sExt.mappingToRectInUIKit(rect)

	fromX := x + width*xOffset
	fromY := y + height*yOffset

	toX := fromX
	toY := fromY + height*distance[0]

	return sExt.s.SwipeFloat(fromX, fromY, toX, toY)
}

func (sExt *SessionExt) SwipeLeft(pathname string, confidence float64, distance ...float64) (err error) {
	return sExt.SwipeLeftOffset(pathname, confidence, 0.9, 0.5, distance...)
}

func (sExt *SessionExt) SwipeLeftOffset(pathname string, confidence, xOffset, yOffset float64, distance ...float64) (err error) {
	if len(distance) == 0 {
		distance = []float64{1.0}
	}
	var rect image.Rectangle
	if rect, err = sExt.findImgRect(pathname, confidence); err != nil {
		return err
	}

	x, y, width, height := sExt.mappingToRectInUIKit(rect)

	fromX := x + width*xOffset
	fromY := y + height*yOffset

	toX := fromX - width*distance[0]
	toY := fromY

	return sExt.s.SwipeFloat(fromX, fromY, toX, toY)
}

func (sExt *SessionExt) SwipeRight(pathname string, confidence float64, distance ...float64) (err error) {
	return sExt.SwipeRightOffset(pathname, confidence, 0.1, 0.5, distance...)
}

func (sExt *SessionExt) SwipeRightOffset(pathname string, confidence, xOffset, yOffset float64, distance ...float64) (err error) {
	if len(distance) == 0 {
		distance = []float64{1.0}
	}
	var rect image.Rectangle
	if rect, err = sExt.findImgRect(pathname, confidence); err != nil {
		return err
	}

	x, y, width, height := sExt.mappingToRectInUIKit(rect)

	fromX := x + width*xOffset
	fromY := y + height*yOffset

	toX := fromX + width*distance[0]
	toY := fromY

	return sExt.s.SwipeFloat(fromX, fromY, toX, toY)
}
