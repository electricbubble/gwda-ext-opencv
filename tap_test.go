package gwda_ext_opencv

import (
	"github.com/electricbubble/gwda"
	"testing"
)

func TestDriverExt_TapWithNumber(t *testing.T) {
	driver, err := gwda.NewUSBDriver(nil)
	checkErr(t, err)

	driverExt, err := Extend(driver, 0.95)
	checkErr(t, err)

	pathSearch := "/Users/hero/Documents/temp/2020-05/opencv/flag7.png"

	// gwda.SetDebug(true)

	err = driverExt.TapWithNumber(pathSearch, 3)
	checkErr(t, err)

	err = driverExt.TapWithNumberOffset(pathSearch, 3, 0.5, 0.75)
	checkErr(t, err)
}
