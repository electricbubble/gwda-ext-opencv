package gwda_ext_opencv

import (
	"github.com/electricbubble/gwda"
	"strconv"
	"strings"
	"testing"
)

func TestSessionExt_GesturePassword(t *testing.T) {
	split := strings.Split("6304258", "")
	password := make([]int, len(split))
	for i := range split {
		password[i], _ = strconv.Atoi(split[i])
	}

	client, err := gwda.NewUSBClient()
	checkErr(t, err)
	session, err := client.NewSession()
	checkErr(t, err)

	sessionExt, err := Extend(session, 0.95)
	checkErr(t, err)

	pathSearch := "/Users/hero/Documents/temp/2020-05/opencv/IMG_5.png"

	err = sessionExt.GesturePassword(pathSearch, password...)
	checkErr(t, err)
}
