package intertitle

import (
	"fmt"

	calibriaImage "github.com/liampulles/cabiria/pkg/image"
	calibriaMath "github.com/liampulles/cabiria/pkg/math"
	"github.com/lucasb-eyer/go-colorful"

	"image"
	"image/color"
)

type IntensityStats struct {
	AvgIntensity       float64
	LowerAvgIntensity  float64
	MiddleAvgIntensity float64
	UpperAvgIntensity  float64
	ProportionLower    float64
	ProportionMiddle   float64
	ProportionUpper    float64
}

func (is IntensityStats) asInput() []float64 {
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

func (is IntensityStats) asCSV() string {
	return fmt.Sprintf("%f,%f,%f,%f,%f,%f,%f",
		is.AvgIntensity, is.LowerAvgIntensity, is.MiddleAvgIntensity,
		is.UpperAvgIntensity, is.ProportionLower, is.ProportionMiddle, is.ProportionUpper)
}

func GetIntensityStats(img image.Image) IntensityStats {
	levelled := calibriaImage.LevelImage(img, 0.1, 0.9)
	sum := float64(0)
	lowerSum := float64(0)
	middleSum := float64(0)
	upperSum := float64(0)
	lowerCount := float64(0.0001) // Avoid divide by zero issues
	middleCount := float64(0.0001)
	upperCount := float64(0.0001)
	calibriaImage.ForEachPixel(levelled, func(x int, y int, col color.Color) {
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

func Intensity(col color.Color) float64 {
	neueCol, _ := colorful.MakeColor(col)
	l, _, _ := neueCol.Luv()
	return calibriaMath.ClampFloat64(l, 0.0, 1.0)
}
