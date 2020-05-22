package main

import (
	. "github.com/electricbubble/gwda"
	extOpenCV "github.com/electricbubble/gwda-ext-opencv"
	"log"
)

func main() {
	client, err := NewClient("http://localhost:8100")
	checkErr(err)
	session, err := client.NewSession()
	checkErr(err)

	sessionExt, err := extOpenCV.Extend(session, "/Users/hero/Documents/temp/2020-05")
	checkErr(err, "扩展 session ，并指定截图保存路径")

	confidence := 0.95

	pathSearch := "/Users/hero/Documents/temp/2020-05/opencv/flag0.png"
	// err = sessionExt.Tap(pathSearch, confidence)
	// checkErr(err, "找到图片(匹配度 >= 95%)后点击（默认 x 向右👉偏移 50%， y 向下👇偏移 50%）")
	//
	// err = sessionExt.TapOffset(pathSearch, confidence, 0.1, 0.1)
	// checkErr(err, "找到图片(匹配度 >= 95%)后点击（ x 向右👉偏移 10%， y 向下👇偏移 10%）")
	//
	pathSearch = "/Users/hero/Documents/temp/2020-05/opencv/flag7.png"
	// err = sessionExt.TapOffset(pathSearch, confidence, 0.2, 0.8)
	// checkErr(err, "找到图片(匹配度 >= 95%)后点击（ x 向右👉偏移 20%， y 向下👇偏移 80%）")

	err = sessionExt.DoubleTap(pathSearch, confidence)
	checkErr(err, "找到图片(匹配度 >= 95%)后双击（默认 x 向右👉偏移 50%， y 向下👇偏移 50%）")

	err = sessionExt.DoubleTapOffset(pathSearch, confidence, 0.1, 0.25)
	checkErr(err, "找到图片(匹配度 >= 95%)后点击（ x 向右👉偏移 10%， y 向下👇偏移 25%）")

	return

	// Debug = true

	pathSearch = "/Users/hero/Documents/temp/2020-05/opencv/flag.png"
	// err = sessionExt.SwipeDown(pathSearch, confidence, 0.5)
	// checkErr(err, "向下👇滑动，滑动距离为图片高度的 50%（默认从图片的正中间顶部向底部滑动，默认滑动距离为 1个 图片高度）")
	//
	// err = sessionExt.SwipeDownOffset(pathSearch, confidence, 0.25, 1)
	// checkErr(err, "向下👇滑动（ x 向右👉偏移 25%， y 向下👇偏移 100% ）")
	//
	// err = sessionExt.SwipeDownOffset(pathSearch, confidence, -0.25, -0.8)
	// checkErr(err, "向下👇滑动（ x 向左👈偏移 25%， y 向上👆偏移 80% ）")
	//
	// return

	// Debug = true

	// err = sessionExt.SwipeUp(pathSearch, confidence, 0.5)
	// checkErr(err, "向上👆滑动，滑动距离为图片高度的 50%（默认从图片的正中间底部向顶部滑动，默认滑动距离为 1个 图片高度）")
	//
	// err = sessionExt.SwipeUpOffset(pathSearch, confidence, 0.9, 0.6)
	// checkErr(err, "向上👆滑动（起始滑动点 x 向右👉偏移 90%， y 向下👇偏移 60% ）")
	//
	// err = sessionExt.SwipeUpOffset(pathSearch, confidence, -0.1, -0.05, 0.3)
	// checkErr(err, "向上👆滑动，滑动距离为图片高度的 30%（起始滑动点 x 向左👈偏移 10%， y 向上👆偏移 5% ）")
	//
	// return

	// err = sessionExt.SwipeLeft(pathSearch, confidence, 0.5)
	// checkErr(err, "向左👈滑动，滑动距离为图片宽度的 50%（默认从图片的正中间右侧向左侧滑动，默认滑动距离为 1个 图片宽度）")
	//
	// err = sessionExt.SwipeLeftOffset(pathSearch, confidence, 0.5, 0.55)
	// checkErr(err, "向左👈滑动（起始滑动点 x 向右👉偏移 50%， y 向下👇偏移 55% ）")
	//
	// err = sessionExt.SwipeLeftOffset(pathSearch, confidence, -0.15, -0.25)
	// checkErr(err, "向左👈滑动（起始滑动点 x 向左👈偏移 15%， y 向上👆偏移 25% ）")
	//
	// return

	err = sessionExt.SwipeRight(pathSearch, confidence, 0.5)
	checkErr(err, "向右👉滑动，滑动距离为图片宽度的 50%（默认从图片的正中间左侧向右侧滑动，默认滑动距离为 1个 图片宽度）")

	err = sessionExt.SwipeRightOffset(pathSearch, confidence, 0.5, 0.6)
	checkErr(err, "向右👉滑动（起始滑动点 x 向右👉偏移 50%， y 向下👇偏移 60% ）")

	err = sessionExt.SwipeRightOffset(pathSearch, 0.92, -0.25, -0.05)
	checkErr(err, "向右👉滑动（起始滑动点 x 向左👈偏移 25%， y 向上👆偏移 5% ）")
}

func checkErr(err error, msg ...string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
