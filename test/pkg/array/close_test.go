package array_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/liampulles/cabiria/pkg/array"
)

func TestCloseBoolArray(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		items     []bool
		threshold uint
		expected  []bool
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
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			array.CloseBoolArray(test.items, test.threshold)

			// Verify result (must be very close)
			if !reflect.DeepEqual(test.items, test.expected) {
				t.Errorf("Unexpected result.\nExpected: %v\nActual: %v", test.expected, test.items)
			}
		})
	}
}

func bools(vals ...int) []bool {
	bools := make([]bool, len(vals))
	for i, val := range vals {
		bools[i] = val > 0
	}
	return bools
}
