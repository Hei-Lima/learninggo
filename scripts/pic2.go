package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	h, w, s    int
	r, g, b, a uint8
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(img.h*img.s, img.w*img.s, 0, 0)
}
func (img Image) At(x, y int) color.Color {
	return color.RGBA{img.r, img.g, img.b, img.a}
}

func main() {
	m := Image{2, 1, 400, 255, 60, 60, 255}
	pic.ShowImage(m)
}
