package image

import (
	"image"
	"image/color"

	"github.com/lucasb-eyer/go-colorful"

	cabiriaMath "github.com/liampulles/cabiria/pkg/math"
	"github.com/liampulles/cabiria/pkg/ml"
	"github.com/liampulles/cabiria/pkg/ml/cluster"
)

const (
	minLightness = 0.3
)

// ChangeValue will convert the color to HSV() space, and return a color
// with V set to value
func ChangeValue(col color.Color, value float64) color.Color {
	clamped := cabiriaMath.ClampFloat64(value, 0.0, 1.0)
	newCol, _ := colorful.MakeColor(col)
	h, s, _ := newCol.Hsv()
	return colorful.Hsv(h, s, clamped)
}

// GetForegroundAndBackground guesses the foreground and background color.
func GetForegroundAndBackground(img image.Image) (color.Color, color.Color, error) {
	// Use KMeans to quantize image into two centroids.
	kMeans := cluster.NewKMeansClassifier(2, 5000)
	counts, _, err := kMeans.Fit(allPixelsAsDatum(img))
	if err != nil {
		return nil, nil, err
	}
	centroids := kMeans.ClusterCentroids()

	// Return least populous centroid for foreground, and the other for background.
	centroidA := datumAsPixel(centroids[0])
	centroidB := datumAsPixel(centroids[1])
	if counts[0] < counts[1] {
		return ChangeValue(centroidA, 1.0), ChangeValue(centroidB, 0.1), nil
	}
	return ChangeValue(centroidB, 1.0), ChangeValue(centroidA, 0.1), nil
}

func allPixelsAsDatum(img image.Image) []ml.Datum {
	result := make([]ml.Datum, img.Bounds().Dx()*img.Bounds().Dy())
	count := 0
	ForEachPixel(img, func(x, y int, col color.Color) {
		result[count] = pixelAsDatum(col)
		count++
	})
	return result
}

func pixelAsDatum(col color.Color) ml.Datum {
	r, g, b, _ := col.RGBA()
	return []float64{float64(r), float64(g), float64(b)}
}

func datumAsPixel(datum ml.Datum) color.Color {
	return color.RGBA{
		R: uint8(datum[0] / 257.0),
		G: uint8(datum[1] / 257.0),
		B: uint8(datum[2] / 257.0),
		A: 255,
	}
}
