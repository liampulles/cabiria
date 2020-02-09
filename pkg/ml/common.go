package ml

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"

	cabiriaMath "github.com/liampulles/cabiria/pkg/math"
)

// Split can be used to split training ang test data by a percentage.
//  The resulting arrays are pseudorandom in composition; not consistent
//  between runs.
func Split(samples []Sample, split float64) ([]Sample, []Sample) {
	rand.Shuffle(len(samples), func(i, j int) { samples[i], samples[j] = samples[j], samples[i] })
	cutoff := int(math.Ceil(float64(len(samples)) * cabiriaMath.ClampFloat64(split, 0.0, 1.0)))
	return samples[:cutoff], samples[cutoff:]
}

// Test runs test samples against a trained predictor to see how accurate it is.
//  The result is percentage accuracy.
func Test(cls Predictor, testData []Sample) (float64, error) {
	passed := 0
	for _, datum := range testData {
		prediction, err := cls.PredictSingle(datum.Input)
		if err != nil {
			return -1.0, err
		}
		match, err := Match(datum.Output, prediction)
		if err != nil {
			return -1.0, err
		}
		if match {
			passed++
		}
	}
	return float64(passed) / float64(len(testData)), nil
}

// Match returns true if the Datum' are the same, otherwise false.
//  An error is returned if the Datum' have different lengths.
func Match(actual Datum, expected Datum) (bool, error) {
	if len(actual) != len(expected) {
		return false, fmt.Errorf("Cannot Match outputs with different lengths. Actual length: %d, expected length: %d",
			len(actual), len(expected))
	}
	for i, actualI := range actual {
		if actualI != expected[i] {
			return false, nil
		}
	}
	return true, nil
}

// AsCSV maps a datum to a CSV line for ML purposes.
func (d Datum) AsCSV() string {
	result := ""
	for i, elem := range d {
		if i > 0 {
			result = result + ","
		}
		result = result + strconv.FormatFloat(elem, 'f', -1, 64)
	}
	return result
}
