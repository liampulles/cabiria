package image

import (
	"fmt"
	"image"
	"image/color"

	"github.com/liampulles/cabiria/pkg/ml"
	"github.com/liampulles/cabiria/pkg/ml/cluster"
)

const (
	minLightness = 0.3
)

// GetForegroundAndBackground guesses the background and foreground color.
func GetForegroundAndBackground(img image.Image) (color.Color, color.Color, error) {
	// Use KMeans to quantize image into two centroids.
	kMeans := cluster.NewKMeansClassifier(2, 5000)
	counts, iter, err := kMeans.Fit(allPixelsAsDatum(img))
	fmt.Printf("%d\n", iter)
	if err != nil {
		return nil, nil, err
	}
	centroids := kMeans.ClusterCentroids()

	// Return least populous centroid for foreground, and the other for background.
	centroidA := datumAsPixel(centroids[0])
	centroidB := datumAsPixel(centroids[1])
	if counts[0] < counts[1] {
		return centroidA, centroidB, nil
	}
	return centroidB, centroidA, nil
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
