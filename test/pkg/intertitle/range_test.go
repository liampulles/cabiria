package intertitle_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/liampulles/cabiria/pkg/intertitle"
)

func TestValid(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		ir       intertitle.Range
		expected bool
	}{
		// Invalid cases
		{
			interRange(-1, 0, 1.0),
			false,
		},
		{
			interRange(0, -1, 1.0),
			false,
		},
		{
			interRange(0, 0, -1.0),
			false,
		},
		{
			interRange(0, 0, 0.0),
			false,
		},
		{
			interRange(1, 0, 1.0),
			false,
		},
		// Valid cases
		{
			interRange(0, 0, 1.0),
			true,
		},
		{
			interRange(1, 2, 2.5),
			true,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual := test.ir.Valid()

			// Verify result
			if actual != test.expected {
				t.Errorf("Result differs. Actual: %v, Expected %v", actual, test.expected)
			}
		})
	}
}

func TestStart(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		ir       intertitle.Range
		expected time.Time
	}{
		{
			interRange(0, 0, 1.0),
			timestamp(0, 0, 0, 0),
		},
		{
			interRange(1, 2, 1.0),
			timestamp(0, 0, 1, 0),
		},
		{
			interRange(10, 20, 2.0),
			timestamp(0, 0, 5, 0),
		},
		{
			interRange(10, 20, 0.5),
			timestamp(0, 0, 20, 0),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual := test.ir.Start()

			// Verify result
			if !actual.Equal(test.expected) {
				t.Errorf("Result differs. Actual: %v, Expected %v", actual, test.expected)
			}
		})
	}
}

func TestEnd(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		ir       intertitle.Range
		expected time.Time
	}{
		{
			interRange(0, 0, 1.0),
			timestamp(0, 0, 0, 0),
		},
		{
			interRange(1, 2, 1.0),
			timestamp(0, 0, 2, 0),
		},
		{
			interRange(10, 20, 2.0),
			timestamp(0, 0, 10, 0),
		},
		{
			interRange(10, 20, 0.5),
			timestamp(0, 0, 40, 0),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual := test.ir.End()

			// Verify result
			if !actual.Equal(test.expected) {
				t.Errorf("Result differs. Actual: %v, Expected %v", actual, test.expected)
			}
		})
	}
}

func TestTransformToNew(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		ir       intertitle.Range
		start    time.Time
		end      time.Time
		expected intertitle.Range
	}{
		{
			interRange(0, 0, 1.0),
			timestamp(0, 0, 0, 0),
			timestamp(0, 0, 0, 0),
			interRange(0, 0, 1.0),
		},
		{
			interRange(1, 2, 1.0),
			timestamp(0, 0, 0, 0),
			timestamp(0, 0, 0, 0),
			interRange(0, 0, 1.0),
		},
		{
			interRange(1, 2, 1.0),
			timestamp(0, 0, 3, 0),
			timestamp(0, 0, 4, 0),
			interRange(3, 4, 1.0),
		},
		{
			interRange(5, 6, 2.0),
			timestamp(0, 0, 2, 0),
			timestamp(0, 0, 4, 0),
			interRange(4, 8, 2.0),
		},
		{
			interRange(5, 6, 0.5),
			timestamp(0, 0, 2, 0),
			timestamp(0, 0, 4, 0),
			interRange(1, 2, 0.5),
		},
		// Truncation case
		{
			interRange(5, 6, 2.5),
			timestamp(0, 0, 3, 0),
			timestamp(0, 0, 4, 0),
			interRange(7, 10, 2.5),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual := test.ir.TransformToNew(test.start, test.end)

			// Verify result
			if actual != test.expected {
				t.Errorf("Result differs. Actual: %v, Expected %v", actual, test.expected)
			}
		})
	}
}

func TestMapRanges(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		intertitles []bool
		fps         float64
		expected    []intertitle.Range
	}{
		// Empty cases
		{
			nil,
			1.0,
			interRanges(),
		},
		{
			intertitles(),
			1.0,
			interRanges(),
		},
		{
			intertitles(0),
			1.0,
			interRanges(),
		},
		{
			intertitles(0, 0, 0, 0),
			1.0,
			interRanges(),
		},
		// A single
		{
			intertitles(1),
			1.0,
			interRanges(
				interRange(0, 0, 1.0),
			),
		},
		{
			intertitles(1, 0),
			1.0,
			interRanges(
				interRange(0, 0, 1.0),
			),
		},
		{
			intertitles(0, 1),
			1.0,
			interRanges(
				interRange(1, 1, 1.0),
			),
		},
		{
			intertitles(0, 1, 0),
			1.0,
			interRanges(
				interRange(1, 1, 1.0),
			),
		},
		// A multiple
		{
			intertitles(1, 1),
			1.0,
			interRanges(
				interRange(0, 1, 1.0),
			),
		},
		{
			intertitles(1, 1, 1),
			1.0,
			interRanges(
				interRange(0, 2, 1.0),
			),
		},
		{
			intertitles(1, 1, 0),
			1.0,
			interRanges(
				interRange(0, 1, 1.0),
			),
		},
		{
			intertitles(0, 1, 1),
			1.0,
			interRanges(
				interRange(1, 2, 1.0),
			),
		},
		{
			intertitles(0, 1, 1, 0),
			1.0,
			interRanges(
				interRange(1, 2, 1.0),
			),
		},
		// Complex case
		{
			intertitles(1, 0, 1, 1, 1, 0, 1, 0, 1, 1),
			1.0,
			interRanges(
				interRange(0, 0, 1.0),
				interRange(2, 4, 1.0),
				interRange(6, 6, 1.0),
				interRange(8, 9, 1.0),
			),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual := intertitle.MapRanges(test.intertitles, test.fps)

			// Verify result
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Result differs. Actual: %v, Expected %v", actual, test.expected)
			}
		})
	}
}

func TestJoinTouchingRanges(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		ir       []intertitle.Range
		expected []intertitle.Range
	}{
		// Empty case
		{
			nil,
			interRanges(),
		},
		{
			interRanges(),
			interRanges(),
		},
		// Single case
		{
			interRanges(
				interRange(1, 5, 1.0),
			),
			interRanges(
				interRange(1, 5, 1.0),
			),
		},
		// Two not touching ranges
		{
			interRanges(
				interRange(1, 5, 1.0),
				interRange(7, 10, 1.0),
			),
			interRanges(
				interRange(1, 5, 1.0),
				interRange(7, 10, 1.0),
			),
		},
		{
			interRanges(
				interRange(7, 10, 1.0),
				interRange(1, 5, 1.0),
			),
			interRanges(
				interRange(1, 5, 1.0),
				interRange(7, 10, 1.0),
			),
		},
		// Two ranges which are one-apart
		{
			interRanges(
				interRange(1, 5, 1.0),
				interRange(6, 10, 1.0),
			),
			interRanges(
				interRange(1, 10, 1.0),
			),
		},
		{
			interRanges(
				interRange(6, 10, 1.0),
				interRange(1, 5, 1.0),
			),
			interRanges(
				interRange(1, 10, 1.0),
			),
		},
		// Two ranges with matching bounds
		{
			interRanges(
				interRange(1, 5, 1.0),
				interRange(5, 10, 1.0),
			),
			interRanges(
				interRange(1, 10, 1.0),
			),
		},
		{
			interRanges(
				interRange(5, 10, 1.0),
				interRange(1, 5, 1.0),
			),
			interRanges(
				interRange(1, 10, 1.0),
			),
		},
		// Two ranges which overlap
		// -> Partial
		{
			interRanges(
				interRange(1, 5, 1.0),
				interRange(3, 10, 1.0),
			),
			interRanges(
				interRange(1, 10, 1.0),
			),
		},
		{
			interRanges(
				interRange(3, 10, 1.0),
				interRange(1, 5, 1.0),
			),
			interRanges(
				interRange(1, 10, 1.0),
			),
		},
		// -> One range inside another
		{
			interRanges(
				interRange(3, 7, 1.0),
				interRange(1, 10, 1.0),
			),
			interRanges(
				interRange(1, 10, 1.0),
			),
		},
		{
			interRanges(
				interRange(1, 10, 1.0),
				interRange(3, 7, 1.0),
			),
			interRanges(
				interRange(1, 10, 1.0),
			),
		},
		// -> Equal ranges
		{
			interRanges(
				interRange(1, 10, 1.0),
				interRange(1, 10, 1.0),
			),
			interRanges(
				interRange(1, 10, 1.0),
			),
		},
		// Three ranges
		{
			interRanges(
				interRange(11, 13, 1.0),
				interRange(3, 7, 1.0),
				interRange(1, 10, 1.0),
			),
			interRanges(
				interRange(1, 13, 1.0),
			),
		},
		// A mixture of cases
		{
			interRanges(
				interRange(18, 21, 1.0),
				interRange(1, 3, 1.0),
				interRange(8, 10, 1.0),
				interRange(18, 22, 1.0),
				interRange(12, 14, 1.0),
				interRange(14, 16, 1.0),
				interRange(20, 24, 1.0),
				interRange(5, 7, 1.0),
				interRange(25, 27, 1.0),
			),
			interRanges(
				interRange(1, 3, 1.0),
				interRange(5, 10, 1.0),
				interRange(12, 16, 1.0),
				interRange(18, 27, 1.0),
			),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual := intertitle.JoinTouchingRanges(test.ir)

			// Verify result
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Result differs. Actual: %v, Expected %v", actual, test.expected)
			}
		})
	}
}

func interRange(start, end int, fps float64) intertitle.Range {
	return intertitle.Range{
		StartFrame: start,
		EndFrame:   end,
		FPS:        fps,
	}
}

func interRanges(interRanges ...intertitle.Range) []intertitle.Range {
	result := make([]intertitle.Range, 0)
	return append(result, interRanges...)
}

func timestamp(hour, min, sec, milli int) time.Time {
	return time.Date(0, time.January, 1, hour, min, sec, milli*1e+6, time.UTC)
}
