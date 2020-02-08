package image

import (
	"image"
	"image/color"
	"math"

	"github.com/liampulles/cabiria/pkg/image/test"

	"github.com/lucasb-eyer/go-colorful"
)

// Diff uses the average euclidiean distance in the CIEL*a*b* colorsace
// to measure image difference.
func Diff(img1 image.Image, img2 image.Image) (float64, error) {
	err := test.ValidateBoundsMatch(img1.Bounds(), img2.Bounds())
	if err != nil {
		return -1.0, err
	}

	total := float64(0)
	ForEachPixel(img1, func(x int, y int, col1 color.Color) {
		col2 := img2.At(x, y)
		total += ColorDiff(col1, col2)
	})
	return total / float64(img1.Bounds().Dx()*img1.Bounds().Dy()), nil
}

// ColorDiff uses the euclidiean distance in the CIEL*a*b* colorsace to measure
// color difference.
func ColorDiff(col1 color.Color, col2 color.Color) float64 {
	neueCol1, _ := colorful.MakeColor(col1)
	neueCol2, _ := colorful.MakeColor(col2)
	l1, a1, b1 := neueCol1.Lab()
	l2, a2, b2 := neueCol2.Lab()
	ldiff := l1 - l2
	adiff := a1 - a2
	bdiff := b1 - b2
	return math.Sqrt((ldiff * ldiff) + (adiff * adiff) + (bdiff * bdiff))
}
