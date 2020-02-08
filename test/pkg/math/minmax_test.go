package math_test

import (
	"fmt"
	"testing"

	"github.com/liampulles/cabiria/pkg/math"
)

func TestMinMaxInt(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		a           int
		b           int
		expectedMin int
		expectedMax int
	}{
		{
			0, 0,
			0, 0,
		},
		{
			0, 1,
			0, 1,
		},
		{
			1, 0,
			0, 1,
		},
		{
			-1, 0,
			-1, 0,
		},
		{
			0, -1,
			-1, 0,
		},
		{
			-2, 5,
			-2, 5,
		},
		{
			5, -2,
			-2, 5,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("(%d,%d) -> (%d,%d)", test.a, test.b, test.expectedMin, test.expectedMax), func(t *testing.T) {
			// Exercise SUT
			actualMin, actualMax := math.MinMaxInt(test.a, test.b)

			// Verify result
			if actualMin != test.expectedMin {
				t.Errorf("Unexpected result.\nExpected Min: %d\nActual Min: %d", test.expectedMin, actualMin)
			}
			if actualMax != test.expectedMax {
				t.Errorf("Unexpected result.\nExpected Max: %d\nActual Max: %d", test.expectedMax, actualMax)
			}
		})
	}
}
