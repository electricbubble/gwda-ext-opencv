package gwda_ext_opencv

import (
	"github.com/electricbubble/gwda"
	"testing"
)

func TestSessionExt_TapWithNumber(t *testing.T) {
	client, err := gwda.NewUSBClient()
	checkErr(t, err)
	session, err := client.NewSession()
	checkErr(t, err)

	sessionExt, err := Extend(session, 0.95)
	checkErr(t, err)

	pathSearch := "/Users/hero/Documents/temp/2020-05/opencv/flag7.png"

	// gwda.WDADebug = true

	err = sessionExt.TapWithNumber(pathSearch, 3)
	checkErr(t, err)

	err = sessionExt.TapWithNumberOffset(pathSearch, 3, 0.5, 0.75)
	checkErr(t, err)
}
