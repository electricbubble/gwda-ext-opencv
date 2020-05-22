# gwda-ext-opencv

[electricbubble/gwda](https://github.com/electricbubble/gwda) 的图片匹配扩展库。

## 安装

> 必须先安装好 `OpenCV`，安装步骤可参考 `hybridgroup/gocv`:
> - [macOS](https://github.com/hybridgroup/gocv#macos) 
> 建议直接用 `Homebrew` 安装

```bash
go get -u github.com/electricbubble/gwda-ext-opencv
```

## 使用
```go
package main

import (
	"fmt"
	. "github.com/electricbubble/gwda"
	extOpenCV "github.com/electricbubble/gwda-ext-opencv"
	"log"
)

func main() {
	client, err := NewClient("http://localhost:8100")
	checkErr(err)
	session, err := client.NewSession()
	checkErr(err)

	sessionExt, err := extOpenCV.Extend(session, 0.95, "/Users/hero/Documents/temp/2020-05")
	checkErr(err, "扩展 session ，指定匹配阀值为 95%（在不修改或者使用 `OnlyOnceThreshold` 的情况下），以及截图保存的路径")

	pathZero := "/Users/hero/Documents/temp/2020-05/opencv/flag0.png"
	err = sessionExt.Tap(pathZero)
	checkErr(err, "找到图片(匹配度 >= 95%)后点击（默认 x 向右👉偏移 50%， y 向下👇偏移 50%）")

	err = sessionExt.TapOffset(pathZero, 0.1, 0.1)
	checkErr(err, "找到图片(匹配度 >= 95%)后点击（ x 向右👉偏移 10%， y 向下👇偏移 10%）")

	pathSeven := "/Users/hero/Documents/temp/2020-05/opencv/flag7.png"
	err = sessionExt.TapOffset(pathSeven, 0.2, 0.8)
	checkErr(err, "找到图片(匹配度 >= 95%)后点击（ x 向右👉偏移 20%， y 向下👇偏移 80%）")

	err = sessionExt.DoubleTap(pathSeven)
	checkErr(err, "找到图片(匹配度 >= 95%)后双击（默认 x 向右👉偏移 50%， y 向下👇偏移 50%）")

	err = sessionExt.DoubleTapOffset(pathSeven, 0.1, 0.25)
	checkErr(err, "找到图片(匹配度 >= 95%)后点击（ x 向右👉偏移 10%， y 向下👇偏移 25%）")

	pathSlash := "/Users/hero/Documents/temp/2020-05/opencv/flag.png"
	err = sessionExt.SwipeDown(pathSlash, 0.5)
	checkErr(err, "向下👇滑动，滑动距离为图片高度的 50%（默认从图片的正中间顶部向底部滑动，默认滑动距离为 1个 图片高度）")

	err = sessionExt.SwipeDownOffset(pathSlash, 0.25, 1)
	checkErr(err, "向下👇滑动（ x 向右👉偏移 25%， y 向下👇偏移 100% ）")

	err = sessionExt.SwipeDownOffset(pathSlash, -0.25, -0.8)
	checkErr(err, "向下👇滑动（ x 向左👈偏移 25%， y 向上👆偏移 80% ）")

	// WDADebug = true

	// 撤销 3次 操作
	undo(sessionExt, 3)

	err = sessionExt.SwipeUp(pathSlash, 0.5)
	checkErr(err, "向上👆滑动，滑动距离为图片高度的 50%（默认从图片的正中间底部向顶部滑动，默认滑动距离为 1个 图片高度）")

	err = sessionExt.SwipeUpOffset(pathSlash, 0.9, 0.6)
	checkErr(err, "向上👆滑动（起始滑动点 x 向右👉偏移 90%， y 向下👇偏移 60% ）")

	err = sessionExt.OnlyOnceThreshold(0.92).SwipeUpOffset(pathSlash, -0.1, -0.05, 0.3)
	checkErr(err, "向上👆滑动，临时指定匹配阀值为 92%，滑动距离为图片高度的 30%（起始滑动点 x 向左👈偏移 10%， y 向上👆偏移 5% ）")

	// 撤销 3次 操作
	undo(sessionExt, 3)

	err = sessionExt.SwipeLeft(pathSlash, 0.5)
	checkErr(err, "向左👈滑动，滑动距离为图片宽度的 50%（默认从图片的正中间右侧向左侧滑动，默认滑动距离为 1个 图片宽度）")

	err = sessionExt.SwipeLeftOffset(pathSlash, 0.5, 0.55)
	checkErr(err, "向左👈滑动（起始滑动点 x 向右👉偏移 50%， y 向下👇偏移 55% ）")

	err = sessionExt.OnlyOnceThreshold(0.92).SwipeLeftOffset(pathSlash, -0.15, -0.25)
	checkErr(err, "向左👈滑动，临时指定匹配阀值为 92%（起始滑动点 x 向左👈偏移 15%， y 向上👆偏移 25% ）")

	// sessionExt.Debug(extOpenCV.DmNotMatch)

	// 撤销 3次 操作
	undo(sessionExt, 3)

	err = sessionExt.SwipeRight(pathSlash, 0.5)
	checkErr(err, "向右👉滑动，滑动距离为图片宽度的 50%（默认从图片的正中间左侧向右侧滑动，默认滑动距离为 1个 图片宽度）")

	err = sessionExt.SwipeRightOffset(pathSlash, 0.5, 0.6)
	checkErr(err, "向右👉滑动（起始滑动点 x 向右👉偏移 50%， y 向下👇偏移 60% ）")

	err = sessionExt.OnlyOnceThreshold(0.90).SwipeRightOffset(pathSlash, -0.25, -0.05)
	checkErr(err, "向右👉滑动（起始滑动点 x 向左👈偏移 25%， y 向上👆偏移 5% ）")

	// 撤销 10次 操作
	undo(sessionExt, 10)
}

func undo(sExt *extOpenCV.SessionExt, n int) {
	pathUndo := "/Users/hero/Documents/temp/2020-05/opencv/undo.png"
	err := sExt.TapWithNumber(pathUndo, n)
	checkErr(err, fmt.Sprintf("撤销 %d次 操作\n", n))
}

func checkErr(err error, msg ...string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

```