package gwda_ext_opencv

func (dExt *DriverExt) Swipe(pathname string, toX, toY int) (err error) {
	return dExt.SwipeFloat(pathname, float64(toX), float64(toY))
}

func (dExt *DriverExt) SwipeFloat(pathname string, toX, toY float64) (err error) {
	return dExt.SwipeOffsetFloat(pathname, toX, toY, 0.5, 0.5)
}

func (dExt *DriverExt) SwipeOffset(pathname string, toX, toY int, xOffset, yOffset float64) (err error) {
	return dExt.SwipeOffsetFloat(pathname, float64(toX), float64(toY), xOffset, yOffset)
}

func (dExt *DriverExt) SwipeOffsetFloat(pathname string, toX, toY, xOffset, yOffset float64) (err error) {
	var x, y, width, height float64
	if x, y, width, height, err = dExt.FindImageRectInUIKit(pathname); err != nil {
		return err
	}

	fromX := x + width*xOffset
	fromY := y + height*yOffset

	return dExt.driver.SwipeFloat(fromX, fromY, toX, toY)
}

func (dExt *DriverExt) SwipeUp(pathname string, distance ...float64) (err error) {
	return dExt.SwipeUpOffset(pathname, 0.5, 0.9, distance...)
}

func (dExt *DriverExt) SwipeUpOffset(pathname string, xOffset, yOffset float64, distance ...float64) (err error) {
	if len(distance) == 0 {
		distance = []float64{1.0}
	}

	var x, y, width, height float64
	if x, y, width, height, err = dExt.FindImageRectInUIKit(pathname); err != nil {
		return err
	}

	fromX := x + width*xOffset
	fromY := (y + height) - height*(1.0-yOffset)

	toX := fromX
	toY := fromY - height*distance[0]

	return dExt.driver.SwipeFloat(fromX, fromY, toX, toY)
}

func (dExt *DriverExt) SwipeDown(pathname string, distance ...float64) (err error) {
	return dExt.SwipeDownOffset(pathname, 0.5, 0.1, distance...)
}

func (dExt *DriverExt) SwipeDownOffset(pathname string, xOffset, yOffset float64, distance ...float64) (err error) {
	if len(distance) == 0 {
		distance = []float64{1.0}
	}

	var x, y, width, height float64
	if x, y, width, height, err = dExt.FindImageRectInUIKit(pathname); err != nil {
		return err
	}

	fromX := x + width*xOffset
	fromY := y + height*yOffset

	toX := fromX
	toY := fromY + height*distance[0]

	return dExt.driver.SwipeFloat(fromX, fromY, toX, toY)
}

func (dExt *DriverExt) SwipeLeft(pathname string, distance ...float64) (err error) {
	return dExt.SwipeLeftOffset(pathname, 0.9, 0.5, distance...)
}

func (dExt *DriverExt) SwipeLeftOffset(pathname string, xOffset, yOffset float64, distance ...float64) (err error) {
	if len(distance) == 0 {
		distance = []float64{1.0}
	}

	var x, y, width, height float64
	if x, y, width, height, err = dExt.FindImageRectInUIKit(pathname); err != nil {
		return err
	}

	fromX := x + width*xOffset
	fromY := y + height*yOffset

	toX := fromX - width*distance[0]
	toY := fromY

	return dExt.driver.SwipeFloat(fromX, fromY, toX, toY)
}

func (dExt *DriverExt) SwipeRight(pathname string, distance ...float64) (err error) {
	return dExt.SwipeRightOffset(pathname, 0.1, 0.5, distance...)
}

func (dExt *DriverExt) SwipeRightOffset(pathname string, xOffset, yOffset float64, distance ...float64) (err error) {
	if len(distance) == 0 {
		distance = []float64{1.0}
	}

	var x, y, width, height float64
	if x, y, width, height, err = dExt.FindImageRectInUIKit(pathname); err != nil {
		return err
	}

	fromX := x + width*xOffset
	fromY := y + height*yOffset

	toX := fromX + width*distance[0]
	toY := fromY

	return dExt.driver.SwipeFloat(fromX, fromY, toX, toY)
}
