package image

import (
	"image"
	"image/color"
	gomath "math"

	"github.com/liampulles/cabiria/pkg/math"

	"github.com/lucasb-eyer/go-colorful"
)

// DetectMaxLum returns the largest pixel Luminance found in the image.
func DetectMaxLum(img image.Image) float64 {
	max := float64(0)
	ForEachPixel(img, func(x int, y int, col color.Color) {
		curr := Lightness(col)
		if curr > max {
			max = curr
		}
	})
	return max
}

// DetectMinLum returns the smallest pixel Luminance found in the image.
func DetectMinLum(img image.Image) float64 {
	min := gomath.MaxFloat64
	ForEachPixel(img, func(x int, y int, col color.Color) {
		curr := Lightness(col)
		if curr < min {
			min = curr
		}
	})
	return min
}

// LevelImage linearly scales the luminance of pixels in an image, such that
// pixels which had luminance min are now black, and pixels which had
// luminance max are now white.
func LevelImage(img image.Image, min float64, max float64) image.Image {
	new := image.NewRGBA(img.Bounds())
	ForEachPixel(img, func(x int, y int, col color.Color) {
		new.Set(x, y, levelColor(col, min, max))
	})
	return new
}

func levelColor(col color.Color, min float64, max float64) color.Color {
	neueCol, _ := colorful.MakeColor(col)
	h, s, l := neueCol.Hsl()
	neueL := levelValue(l, min, max)
	return colorful.Hsl(h, s, neueL)
}

func levelValue(val float64, min float64, max float64) float64 {
	return math.ClampFloat64((val-min)/(max-min), 0.0, 1.0)
}

// Lightness uses L in HSL.
func Lightness(col color.Color) float64 {
	// Convert to colorful color
	neueCol, _ := colorful.MakeColor(col)
	_, _, l := neueCol.Hsl()
	return math.ClampFloat64(l, 0.0, 1.0)
}
