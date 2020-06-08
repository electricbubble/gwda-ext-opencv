package gwda_ext_opencv

import (
	"github.com/electricbubble/gwda"
	"testing"
)

func TestSessionExt_Drag(t *testing.T) {
	client, err := gwda.NewUSBClient()
	checkErr(t, err)
	session, err := client.NewSession()
	checkErr(t, err)

	sessionExt, err := Extend(session, 0.95)
	checkErr(t, err)

	pathSearch := "/Users/hero/Documents/temp/2020-05/opencv/IMG_map.png"

	// err = sessionExt.Drag(pathSearch, 300, 500, 2)
	// checkErr(t, err)

	err = sessionExt.DragOffset(pathSearch, 300, 500, 2.1, 0.5, 2)
	checkErr(t, err)
}
