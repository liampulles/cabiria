package image

import (
	"image"
	"image/color"
)

// parameters: x, y, pixel.
type pixelFunc func(int, int, color.Color)

// ForEachPixel performs pixelFunc on every pixel in the image, going left
// to right, top to bottom.
func ForEachPixel(img image.Image, f pixelFunc) {
	v, ok := img.(*image.RGBA)
	if ok {
		fastForEachPixelRGBA(v, f)
	} else {
		slowForEachPixel(img, f)
	}
}

func slowForEachPixel(img image.Image, f pixelFunc) {
	b := img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			f(x, y, img.At(x, y))
		}
	}
}

func fastForEachPixelRGBA(img *image.RGBA, f pixelFunc) {
	b := img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			f(x, y, img.RGBAAt(x, y))
		}
	}
}
