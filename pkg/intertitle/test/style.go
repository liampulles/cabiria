package test

import (
	"fmt"

	"github.com/liampulles/cabiria/pkg/image/test"

	"github.com/liampulles/cabiria/pkg/intertitle"
)

// CompareStyle will return an error if actual and expected differ.
func CompareStyle(actual, expected intertitle.Style) error {
	if actual.ForegroundColor == nil || expected.ForegroundColor == nil {
		if expected.ForegroundColor == nil && actual.ForegroundColor == nil {
			return nil
		}
		return fmt.Errorf("nil value: Actual Foreground: %v, Expected Foreground %v", actual, expected)
	}
	if actual.BackgroundColor == nil || expected.BackgroundColor == nil {
		if expected.BackgroundColor == nil && actual.BackgroundColor == nil {
			return nil
		}
		return fmt.Errorf("nil value: Actual Background: %v, Expected Background %v", actual, expected)
	}

	if err := test.CompareColor(actual.ForegroundColor, expected.ForegroundColor); err != nil {
		return fmt.Errorf("Foregrounds differ: %v", err)
	}
	if err := test.CompareColor(actual.BackgroundColor, expected.BackgroundColor); err != nil {
		return fmt.Errorf("Backgrounds differ: %v", err)
	}
	return nil
}
