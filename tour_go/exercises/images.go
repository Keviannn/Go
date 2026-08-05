package main

import (
    "image"
    "image/color"
    "golang.org/x/tour/pic"
)

type Image struct{}

func (Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, 256, 256)
}

// Color generation is fake, does not refer to a real image
func (Image) At(x, y int) color.Color {
    v := uint8(x ^ y)
    return color.RGBA{v, v*2, v*3, 255}
}

func main() {
    m := Image{}
    pic.ShowImage(m)
}
