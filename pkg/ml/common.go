package ml

import (
	"fmt"
	"math"
	"math/rand"

	cabiriaMath "github.com/liampulles/cabiria/pkg/math"
)

func Split(samples []Sample, split float64) ([]Sample, []Sample) {
	rand.Shuffle(len(samples), func(i, j int) { samples[i], samples[j] = samples[j], samples[i] })
	cutoff := int(math.Ceil(float64(len(samples)) * cabiriaMath.ClampFloat64(split, 0.0, 1.0)))
	return samples[:cutoff], samples[cutoff:]
}

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

func Match(actual Datum, expected Datum) (bool, error) {
	if len(actual) != len(expected) {
		return false, fmt.Errorf("Cannot Match outputs with different lengths. Actual length: %d, expected length: %d",
			len(actual), len(expected))
	}
	for i, actuali := range actual {
		if actuali != expected[i] {
			return false, nil
		}
	}
	return true, nil
}
