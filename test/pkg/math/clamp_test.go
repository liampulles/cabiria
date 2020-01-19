package math_test

import (
	"fmt"
	"testing"

	"github.com/liampulles/cabiria/pkg/math"
)

func TestClamp(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		min      float64
		max      float64
		val      float64
		expected float64
	}{
		{
			0.0,
			0.0,
			0.0,
			0.0,
		},
		{
			0.0,
			1.0,
			0.0,
			0.0,
		},
		{
			1.0,
			1.0,
			0.0,
			1.0,
		},
		{
			0.0,
			1.0,
			0.5,
			0.5,
		},
		{
			0.0,
			1.0,
			-0.5,
			0.0,
		},
		{
			0.0,
			1.0,
			1.5,
			1.0,
		},
		{
			-1.0,
			0.0,
			-0.5,
			-0.5,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%f -> [%f,%f] -> %f", test.val, test.min, test.max, test.expected), func(t *testing.T) {
			// Exercise SUT
			actual := math.Clamp(test.val, test.min, test.max)

			// Verify result
			if actual != test.expected {
				t.Errorf("Unexpected result.\nExpected: %f\nActual: %f", test.expected, actual)
			}
		})
	}
}
