package array_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/liampulles/cabiria/pkg/array"
)

func TestCloseBoolArrayAndOpenBoolArray(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		items     []bool
		threshold uint
		expected  []bool // This is for close, open expected is the inversion.
	}{
		// Empty
		{
			bools(),
			0,
			bools(),
		},
		{
			bools(),
			1,
			bools(),
		},
		{
			bools(),
			5,
			bools(),
		},
		// Single off
		{
			bools(0),
			0,
			bools(0),
		},
		{
			bools(0),
			1,
			bools(0),
		},
		{
			bools(0),
			5,
			bools(0),
		},
		// Single on
		{
			bools(1),
			0,
			bools(1),
		},
		{
			bools(1),
			1,
			bools(1),
		},
		{
			bools(1),
			5,
			bools(1),
		},
		// Two off
		{
			bools(0, 0),
			0,
			bools(0, 0),
		},
		{
			bools(0, 0),
			1,
			bools(0, 0),
		},
		{
			bools(0, 0),
			5,
			bools(0, 0),
		},
		// Two on
		{
			bools(1, 1),
			0,
			bools(1, 1),
		},
		{
			bools(1, 1),
			1,
			bools(1, 1),
		},
		{
			bools(1, 1),
			5,
			bools(1, 1),
		},
		// One on, one off
		{
			bools(0, 1),
			0,
			bools(0, 1),
		},
		{
			bools(0, 1),
			1,
			bools(0, 1),
		},
		{
			bools(0, 1),
			5,
			bools(0, 1),
		},
		{
			bools(1, 0),
			0,
			bools(1, 0),
		},
		{
			bools(1, 0),
			1,
			bools(1, 0),
		},
		{
			bools(1, 0),
			5,
			bools(1, 0),
		},
		// Two off, and one on in the middle.
		{
			bools(0, 1, 0),
			0,
			bools(0, 1, 0),
		},
		{
			bools(0, 1, 0),
			1,
			bools(0, 1, 0),
		},
		{
			bools(0, 1, 0),
			5,
			bools(0, 1, 0),
		},
		// Two on, one off in the middle
		{
			bools(1, 0, 1),
			0,
			bools(1, 0, 1),
		},
		{
			bools(1, 0, 1),
			1,
			bools(1, 0, 1),
		},
		{
			bools(1, 0, 1),
			2,
			bools(1, 1, 1),
		},
		{
			bools(1, 0, 1),
			5,
			bools(1, 1, 1),
		},
		// Several on, one off in the middle.
		{
			bools(1, 1, 0, 1),
			0,
			bools(1, 1, 0, 1),
		},
		{
			bools(1, 1, 0, 1),
			1,
			bools(1, 1, 0, 1),
		},
		{
			bools(1, 1, 0, 1),
			2,
			bools(1, 1, 1, 1),
		},
		{
			bools(1, 1, 0, 1),
			5,
			bools(1, 1, 1, 1),
		},
		// Two on, several off in the middle.
		{
			bools(1, 0, 0, 1),
			0,
			bools(1, 0, 0, 1),
		},
		{
			bools(1, 0, 0, 1),
			1,
			bools(1, 0, 0, 1),
		},
		{
			bools(1, 0, 0, 1),
			2,
			bools(1, 0, 0, 1),
		},
		{
			bools(1, 0, 0, 1),
			3,
			bools(1, 1, 1, 1),
		},
		{
			bools(1, 0, 0, 1),
			5,
			bools(1, 1, 1, 1),
		},
		// Complex case
		{
			bools(1, 1, 0, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1, 0),
			0,
			bools(1, 1, 0, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1, 0),
		},
		{
			bools(1, 1, 0, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1, 0),
			1,
			bools(1, 1, 0, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1, 0),
		},
		{
			bools(1, 1, 0, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1, 0),
			2,
			bools(1, 1, 1, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1, 0),
		},
		{
			bools(1, 1, 0, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1, 0),
			3,
			bools(1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0),
		},
		{
			bools(1, 1, 0, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1, 0),
			4,
			bools(1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0),
		},
		{
			bools(1, 1, 0, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1, 0),
			5,
			bools(1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0),
		},
	}

	for i, test := range tests {
		closeInput := make([]bool, len(test.items))
		copy(closeInput, test.items)
		closeExpected := make([]bool, len(test.expected))
		copy(closeExpected, test.expected)
		t.Run(fmt.Sprintf("[%d]: Close", i), func(t *testing.T) {
			// Exercise SUT
			array.CloseBoolArray(closeInput, test.threshold)

			// Verify result (must be very close)
			if !reflect.DeepEqual(closeInput, closeExpected) {
				t.Errorf("Unexpected result.\nExpected: %v\nActual: %v", closeExpected, closeInput)
			}
		})

		openInput := make([]bool, len(test.items))
		copy(openInput, test.items)
		openInput = invert(openInput)
		openExpected := make([]bool, len(test.expected))
		copy(openExpected, test.expected)
		openExpected = invert(openExpected)
		t.Run(fmt.Sprintf("[%d]: Open", i), func(t *testing.T) {
			// Exercise SUT
			array.OpenBoolArray(openInput, test.threshold)

			// Verify result (must be very close)
			if !reflect.DeepEqual(openInput, openExpected) {
				t.Errorf("Unexpected result.\nExpected: %v\nActual: %v", openExpected, openInput)
			}
		})
	}
}

func bools(items ...int) []bool {
	bools := make([]bool, len(items))
	for i, item := range items {
		bools[i] = item > 0
	}
	return bools
}

func invert(items []bool) []bool {
	result := make([]bool, len(items))
	for i, elem := range items {
		result[i] = !elem
	}
	return result
}
