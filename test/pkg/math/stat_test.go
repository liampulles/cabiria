package math_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/liampulles/cabiria/pkg/math"
)

func TestAdd(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		a        []float64
		b        []float64
		expected []float64
	}{
		{
			nil,
			nil,
			[]float64{},
		},
		{
			nil,
			[]float64{},
			[]float64{},
		},
		{
			[]float64{},
			[]float64{},
			[]float64{},
		},
		{
			[]float64{0},
			[]float64{0},
			[]float64{0},
		},
		{
			[]float64{0},
			[]float64{1},
			[]float64{1},
		},
		{
			[]float64{0, 10, 100},
			[]float64{-1, 2, -30},
			[]float64{-1, 12, 70},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual, err := math.Add(test.a, test.b)

			// Verify result
			if err != nil {
				t.Errorf("SUT returned an error: %v", err)
			}
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Unexpected result: Expected: %v, Actual: %v", test.expected, actual)
			}
		})
	}
}

func TestAdd_WhenInputArraysDifferInSize(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		a []float64
		b []float64
	}{
		{
			nil,
			[]float64{0},
		},
		{
			[]float64{0},
			nil,
		},
		{
			[]float64{},
			[]float64{0},
		},
		{
			[]float64{0},
			[]float64{},
		},
		{
			[]float64{0},
			[]float64{1, 2},
		},
		{
			[]float64{1, 2},
			[]float64{0},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			_, err := math.Add(test.a, test.b)

			// Verify result
			if err == nil {
				t.Errorf("Expected SUT to return an error.")
			}
		})
	}
}

func TestSquare(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		a        []float64
		expected []float64
	}{
		{
			nil,
			[]float64{},
		},
		{
			[]float64{},
			[]float64{},
		},
		{
			[]float64{0},
			[]float64{0},
		},
		{
			[]float64{-1},
			[]float64{1},
		},
		{
			[]float64{1},
			[]float64{1},
		},
		{
			[]float64{-2, -1, 0, 1, 2},
			[]float64{4, 1, 0, 1, 4},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual := math.Square(test.a)

			// Verify result
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Unexpected result: Expected: %v, Actual: %v", test.expected, actual)
			}
		})
	}
}
