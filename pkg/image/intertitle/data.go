package intertitle

import (
	"github.com/liampulles/cabiria/pkg/ml"

	cabiriaImage "github.com/liampulles/cabiria/pkg/image"
	cabiriaMath "github.com/liampulles/cabiria/pkg/math"
	"github.com/lucasb-eyer/go-colorful"

	"image"
	"image/color"
)

// IntensityStats defines statistics about an image's luminosity
//  in various capacities
type IntensityStats struct {
	AvgIntensity       float64
	LowerAvgIntensity  float64
	MiddleAvgIntensity float64
	UpperAvgIntensity  float64
	ProportionLower    float64
	ProportionMiddle   float64
	ProportionUpper    float64
}

// AsInput maps intensity statistics to a Datum for ML purposes
func (is IntensityStats) AsInput() ml.Datum {
	return []float64{
		is.AvgIntensity,
		is.LowerAvgIntensity,
		is.MiddleAvgIntensity,
		is.UpperAvgIntensity,
		is.ProportionLower,
		is.ProportionMiddle,
		is.ProportionUpper,
	}
}

// GetIntensityStats extracts intensity statistics from an image.
func GetIntensityStats(img image.Image) IntensityStats {
	levelled := cabiriaImage.LevelImage(img, 0.1, 0.9)
	sum := float64(0)
	lowerSum := float64(0)
	middleSum := float64(0)
	upperSum := float64(0)
	lowerCount := float64(0.0001) // Avoid divide by zero issues
	middleCount := float64(0.0001)
	upperCount := float64(0.0001)
	cabiriaImage.ForEachPixel(levelled, func(x int, y int, col color.Color) {
		intensity := Intensity(col)
		sum += intensity
		if intensity < 0.25 {
			lowerSum += intensity
			lowerCount++
		} else if intensity < 0.75 {
			middleSum += intensity
			middleCount++
		} else {
			upperSum += intensity
			upperCount++
		}
	})
	count := lowerCount + middleCount + upperCount
	return IntensityStats{
		AvgIntensity:       sum / count,
		LowerAvgIntensity:  lowerSum / lowerCount,
		MiddleAvgIntensity: middleSum / middleCount,
		UpperAvgIntensity:  upperSum / upperCount,
		ProportionLower:    lowerCount / count,
		ProportionMiddle:   middleCount / count,
		ProportionUpper:    upperCount / count,
	}
}

// Intensity retrieves the L component of the Luv transformation of col.
func Intensity(col color.Color) float64 {
	newCol, _ := colorful.MakeColor(col)
	l, _, _ := newCol.Luv()
	return cabiriaMath.ClampFloat64(l, 0.0, 1.0)
}
