package gwda_ext_opencv

import (
	"github.com/electricbubble/gwda"
	"testing"
)

func TestSessionExt_Swipe(t *testing.T) {
	client, err := gwda.NewClient("http://localhost:8100")
	checkErr(t, err)
	session, err := client.NewSession()
	checkErr(t, err)

	sessionExt, err := Extend(session, 0.95, "/Users/hero/Documents/temp/2020-05")
	checkErr(t, err)

	pathSearch := "/Users/hero/Documents/temp/2020-05/opencv/flag7.png"

	// gwda.WDADebug = true

	// err = sessionExt.Swipe(pathSearch, 300, 500)
	// checkErr(t, err)
	//
	// err = sessionExt.SwipeFloat(pathSearch, 300.9, 500)
	// checkErr(t, err)

	// err = sessionExt.SwipeOffset(pathSearch, 300, 500, 0.2, 0.5)
	// checkErr(t, err)

	// sessionExt.Debug(DmNotMatch)

	err = sessionExt.OnlyOnceThreshold(0.92).SwipeOffsetFloat(pathSearch, 300.9, 499.1, 0.2, 0.5)
	checkErr(t, err)
}
