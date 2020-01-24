package array_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/liampulles/cabiria/pkg/array"
)

type morphFunc func([]bool, uint) []bool

func TestMorphologicalOperations(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		set              []bool
		kernel           uint
		expectedDilation []bool
		expectedErosion  []bool
		expectedClosing  []bool
		expectedOpening  []bool
	}{
		{
			[]bool{},
			0,
			[]bool{},
			[]bool{},
			[]bool{},
			[]bool{},
		},
		{
			[]bool{},
			100,
			[]bool{},
			[]bool{},
			[]bool{},
			[]bool{},
		},
		{
			set(0),
			0,
			set(0),
			set(0),
			set(0),
			set(0),
		},
		{
			set(0),
			1,
			set(0),
			set(0),
			set(0),
			set(0),
		},
		{
			set(0),
			100,
			set(0),
			set(0),
			set(0),
			set(0),
		},
		{
			set(1),
			0,
			set(1),
			set(1),
			set(1),
			set(1),
		},
		{
			set(1),
			1,
			set(1),
			set(1),
			set(1),
			set(1),
		},
		{
			set(1),
			100,
			set(1),
			set(0),
			set(0),
			set(0),
		},
		{
			set(0, 0),
			0,
			set(0, 0),
			set(0, 0),
			set(0, 0),
			set(0, 0),
		},
		{
			set(0, 0),
			1,
			set(0, 0),
			set(0, 0),
			set(0, 0),
			set(0, 0),
		},
		{
			set(0, 0),
			100,
			set(0, 0),
			set(0, 0),
			set(0, 0),
			set(0, 0),
		},
		{
			set(1, 0),
			0,
			set(1, 0),
			set(1, 0),
			set(1, 0),
			set(1, 0),
		},
		{
			set(1, 0),
			1,
			set(1, 0),
			set(1, 0),
			set(1, 0),
			set(1, 0),
		},
		{
			set(1, 0),
			2,
			set(1, 1),
			set(0, 0),
			set(0, 0),
			set(0, 1),
		},
		{
			set(1, 0),
			2,
			set(1, 1),
			set(1, 0),
			set(1, 1),
			set(1, 1),
		},
		{
			set(0, 0, 0),
			0,
			set(0, 0, 0),
			set(0, 0, 0),
			set(0, 0, 0),
			set(0, 0, 0),
		},
		{
			set(0, 0, 0),
			1,
			set(0, 0, 0),
			set(0, 0, 0),
			set(0, 0, 0),
			set(0, 0, 0),
		},
		{
			set(0, 0, 0),
			100,
			set(0, 0, 0),
			set(0, 0, 0),
			set(0, 0, 0),
			set(0, 0, 0),
		},
		{
			set(1, 0, 0),
			0,
			set(1, 0, 0),
			set(1, 0, 0),
			set(1, 0, 0),
			set(1, 0, 0),
		},
		{
			set(1, 0, 0),
			1,
			set(1, 0, 0),
			set(1, 0, 0),
			set(1, 0, 0),
			set(1, 0, 0),
		},
		{
			set(1, 0, 0),
			2,
			set(1, 1, 0),
			set(1, 0, 0),
			set(1, 1, 0),
			set(1, 0, 0),
		},
	}

	for i, test := range tests {
		var suts = []struct {
			name        string
			sut         morphFunc
			expected    []bool
			idempotency bool
		}{
			{
				"Dilation",
				array.Dilation,
				test.expectedDilation,
				false,
			},
			{
				"Erosion",
				array.Erosion,
				test.expectedErosion,
				false,
			},
			{
				"Closing",
				array.Closing,
				test.expectedClosing,
				true,
			},
			{
				"Opening",
				array.Opening,
				test.expectedOpening,
				true,
			},
		}
		for _, sut := range suts {
			t.Run(fmt.Sprintf("[%d/%d: %s]", i+1, len(tests), sut.name), func(t *testing.T) {

				// Exercise SUTs
				actual := sut.sut(test.set, test.kernel)

				// Verify results
				if !reflect.DeepEqual(actual, sut.expected) {
					t.Errorf("Unexpected %s result.\nExpected: %v\nActual: %v", sut.name, sut.expected, actual)
				}

				if sut.idempotency {
					// Verify idempotency
					actual2 := sut.sut(actual, test.kernel)
					if !reflect.DeepEqual(actual, actual2) {
						t.Errorf("%s result is not idempotent.\nActual: %v\nActual 2: %v", sut.name, actual, actual2)
					}
				}
			})
		}
	}
}

func set(items ...int) []bool {
	set := make([]bool, len(items))
	for i, item := range items {
		if item == 1 {
			set[i] = true
		} else {
			set[i] = false
		}
	}
	return set
}
