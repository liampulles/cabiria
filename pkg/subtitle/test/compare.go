package test

import (
	"fmt"

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
	if !actual.Start.Equal(expected.Start) {
		return fmt.Errorf("start differs: Actual: %s, Expected: %s", actual.Start, expected.Start)
	}
	if !actual.End.Equal(expected.End) {
		return fmt.Errorf("end differs: Actual: %s, Expected: %s", actual.End, expected.End)
	}
	return nil
}
