package test

import (
	"fmt"
	"image"
)

// CompareImage will return an error if there is any difference in the colors
// of corresponding pixels in actual and expected.
func CompareImage(actual image.Image, expected image.Image) error {
	actualB := actual.Bounds()
	expectedB := expected.Bounds()
	err := ValidateBoundsMatch(actualB, expectedB)
	if err != nil {
		return err
	}

	for y := actualB.Min.Y; y < actualB.Max.Y; y++ {
		for x := actualB.Min.X; x < actualB.Max.X; x++ {
			ra, ba, ga, aa := actual.At(x, y).RGBA()
			re, be, ge, ae := expected.At(x, y).RGBA()
			if ra != re || ba != be || ga != ge || aa != ae {
				return fmt.Errorf("Color mismatch at (%d,%d).\n"+
					"Actual RGBA: (%d,%d,%d,%d)\n"+
					"Expected RGBA: (%d,%d,%d,%d)",
					x, y,
					ra, ga, ba, aa,
					re, ge, be, ae)
			}
		}
	}
	return nil
}

// ValidateBoundsMatch will return an error if the width or height of
// actualB and expectedB differ.
func ValidateBoundsMatch(actualB image.Rectangle, expectedB image.Rectangle) error {

	if actualB.Size().X != expectedB.Dx() {
		return fmt.Errorf("Actual and expected have different widths.\n"+
			"Actual width: %d\n"+
			"Expected width: %d",
			actualB.Dx(), expectedB.Dx())
	}
	if actualB.Dy() != expectedB.Dy() {
		return fmt.Errorf("Actual and expected have different heights.\n"+
			"Actual height: %d\n"+
			"Expected height: %d",
			actualB.Dy(), expectedB.Dy())
	}
	return nil
}
