package math_test

import (
	"fmt"
	"math"
	"testing"

	calibriaMath "github.com/liampulles/cabiria/pkg/math"
)

func TestEuclideanDistance_ValidInput_ShouldPass(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		a        []float64
		b        []float64
		expected float64
	}{
		{
			[]float64{},
			[]float64{},
			0.0,
		},
		{
			[]float64{0.0},
			[]float64{0.0},
			0.0,
		},
		{
			[]float64{0.0},
			[]float64{1.0},
			1.0,
		},
		{
			[]float64{0.0, 0.0},
			[]float64{1.0, 1.0},
			math.Sqrt(2.0),
		},
		{
			[]float64{-1.0, 2.0},
			[]float64{2.0, -2.0},
			5.0,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("(%v,%v) -> %f", test.a, test.b, test.expected), func(t *testing.T) {
			// Exercise SUT
			actual, err := calibriaMath.EuclideanDistance(test.a, test.b)

			// Verify result
			if err != nil {
				t.Errorf("Encountered error while executing SUT: %v", err)
			}
			if actual != test.expected {
				t.Errorf("Unexpected result.\nExpected: %f\nActual: %f", test.expected, actual)
			}
		})
	}
}

func TestEuclideanDistance_InputWithDifferentLength_ShouldFail(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		a []float64
		b []float64
	}{
		{
			nil,
			nil,
		},
		{
			[]float64{0.0},
			nil,
		},
		{
			nil,
			[]float64{0.0},
		},
		{
			[]float64{},
			[]float64{0.0},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("(%v,%v)", test.a, test.b), func(t *testing.T) {
			// Exercise SUT
			_, err := calibriaMath.EuclideanDistance(test.a, test.b)

			// Verify result
			if err == nil {
				t.Errorf("Expected SUT to throw an error")
			}
		})
	}
}
