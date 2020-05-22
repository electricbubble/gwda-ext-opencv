package gwda_ext_opencv

import (
	"github.com/electricbubble/gwda"
	"testing"
)

func TestSessionExt_TapWithNumber(t *testing.T) {
	client, err := gwda.NewClient("http://localhost:8100")
	checkErr(t, err)
	session, err := client.NewSession()
	checkErr(t, err)

	sessionExt, err := Extend(session, 0.95, "/Users/hero/Documents/temp/2020-05")
	checkErr(t, err, "扩展 session ，并指定截图保存路径")

	pathSearch := "/Users/hero/Documents/temp/2020-05/opencv/flag7.png"

	// gwda.WDADebug = true

	err = sessionExt.TapWithNumber(pathSearch, 3)
	checkErr(t, err)

	err = sessionExt.TapWithNumberOffset(pathSearch, 3, 0.5, 0.75)
	checkErr(t, err)
}
