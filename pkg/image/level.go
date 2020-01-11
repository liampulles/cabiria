package image

import (
	"image"
	"image/color"
	gomath "math"

	"github.com/lucasb-eyer/go-colorful"
	_ "github.com/lucasb-eyer/go-colorful"
)

func DetectMaxLum(img image.Image) float64 {
	max := float64(0)
	ForEachPixel(img, func(x int, y int, col color.Color) {
		curr := Luminance(col)
		if curr > max {
			max = curr
		}
	})
	return max
}

func DetectMinLum(img image.Image) float64 {
	min := gomath.MaxFloat64
	ForEachPixel(img, func(x int, y int, col color.Color) {
		curr := Luminance(col)
		if curr < min {
			min = curr
		}
	})
	return min
}

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
	return clamp((val-min)/(max-min), 0.0, 1.0)

}

func Luminance(col color.Color) float64 {
	// Convert to colorful color
	neueCol, _ := colorful.MakeColor(col)
	_, _, l := neueCol.Hsl()
	return clamp(l, 0.0, 1.0)
}

func clamp(val float64, min float64, max float64) float64 {
	if val <= min {
		return min
	}
	if val >= max {
		return max
	}
	return val
}
