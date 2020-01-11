package test

import (
	"fmt"
	"image"
)

func CompareImage(actual image.Image, expected image.Image) error {
	actualB := actual.Bounds()
	expectedB := expected.Bounds()
	if actualB.Size().X != expectedB.Size().X {
		return fmt.Errorf("Actual and expected have different widths.\n"+
			"Actual width: %d\n"+
			"Expected width: %d",
			actualB.Size().X, expectedB.Size().X)
	}
	if actualB.Size().Y != expectedB.Size().Y {
		return fmt.Errorf("Actual and expected have different heights.\n"+
			"Actual height: %d\n"+
			"Expected height: %d",
			actualB.Size().Y, expectedB.Size().Y)
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
