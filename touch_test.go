package gwda_ext_opencv

import (
	"github.com/electricbubble/gwda"
	"testing"
)

func TestSessionExt_ForceTouch(t *testing.T) {
	client, err := gwda.NewUSBClient()
	checkErr(t, err)
	session, err := client.NewSession()
	checkErr(t, err)

	sessionExt, err := Extend(session, 0.95)
	checkErr(t, err)

	pathSearch := "/Users/hero/Documents/temp/2020-05/opencv/IMG_ft.png"

	err = sessionExt.ForceTouch(pathSearch, 0.5, 3)
	checkErr(t, err)

	// err = sessionExt.ForceTouchOffset(pathSearch, 0.5, 0.1, 0.9)
	// checkErr(t, err)

	// err = sessionExt.ForceTouchOffset(pathSearch, 0.2, 1.1, -1)
	// checkErr(t, err)
}

func TestSessionExt_TouchAndHold(t *testing.T) {
	client, err := gwda.NewUSBClient()
	checkErr(t, err)
	session, err := client.NewSession()
	checkErr(t, err)

	sessionExt, err := Extend(session, 0.95)
	checkErr(t, err)

	pathSearch := "/Users/hero/Documents/temp/2020-05/opencv/IMG_ft.png"

	// err = sessionExt.TouchAndHold(pathSearch)
	// checkErr(t, err)

	// err = sessionExt.TouchAndHold(pathSearch, 3)
	// checkErr(t, err)

	err = sessionExt.TouchAndHoldOffset(pathSearch, 0.8, 0.1)
	checkErr(t, err)
}

func checkErr(t *testing.T, err error, msg ...string) {
	if err != nil {
		if len(msg) == 0 {
			t.Fatal(err)
		} else {
			t.Fatal(msg, err)
		}
	}
}
