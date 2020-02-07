package test

import (
	"fmt"
	"math"
	"time"

	"github.com/liampulles/cabiria/pkg/subtitle"
)

func CompareSubtitles(actual, expected []subtitle.Subtitle) error {
	if actual == nil || expected == nil {
		if expected == nil && actual == nil {
			return nil
		}
		return fmt.Errorf("nil value: Actual: %v, Expected %v", actual, expected)
	}

	if len(actual) != len(expected) {
		return fmt.Errorf("different lengths: Actual: %v, Expected %v", len(actual), len(expected))
	}

	for i, actualI := range actual {
		expectedI := expected[i]
		err := CompareSubtitle(actualI, expectedI)
		if err != nil {
			return fmt.Errorf("comparison failure on element %d: %v", i, err)
		}
	}

	return nil
}

func CompareSubtitle(actual, expected subtitle.Subtitle) error {
	if actual.Text != expected.Text {
		return fmt.Errorf("text differs: Actual: %s, Expected: %s", actual.Text, expected.Text)
	}
	if !veryClose(actual.StartTime, expected.StartTime) {
		return fmt.Errorf("startTime differs: Actual: %s, Expected: %s", actual.StartTime, expected.StartTime)
	}
	if !veryClose(actual.EndTime, expected.EndTime) {
		return fmt.Errorf("endTime differs: Actual: %s, Expected: %s", actual.EndTime, expected.EndTime)
	}
	return nil
}

func veryClose(actual, expected time.Time) bool {
	return math.Abs(float64(actual.Sub(expected))) < float64(50*time.Nanosecond)
}
