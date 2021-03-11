package main

import "gopkg.in/gographics/imagick.v3/imagick"

func main() {
	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	mw.ReadImage("test.pdf")
	mw.SetIteratorIndex(0) // This being the page offset
	mw.SetImageFormat("jpg")
	mw.WriteImage("test.jpg")
}
