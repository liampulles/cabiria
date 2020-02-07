package periods_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/liampulles/cabiria/pkg/time/period"
)

func TestDoesOverlap(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		a        period.Period
		b        period.Period
		expected bool
	}{
		// nil cases
		{
			nil,
			nil,
			false,
		},
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			nil,
			false,
		},
		{
			nil,
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			false,
		},
		// Completely seperate
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			testPeriod{1, timestamp(0, 0, 3, 0), timestamp(0, 0, 4, 0)},
			false,
		},
		{
			testPeriod{1, timestamp(0, 0, 3, 0), timestamp(0, 0, 4, 0)},
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			false,
		},
		// Edges touching
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
			false,
		},
		{
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			false,
		},
		// Partial overlap
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 5, 0)},
			true,
		},
		{
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 5, 0)},
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
			true,
		},
		// One period contained in another
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 5, 0)},
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
			true,
		},
		{
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 5, 0)},
			true,
		},
		// Equal periods
		{
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
			true,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual := period.DoesOverlap(test.a, test.b)

			// Verify result
			if actual != test.expected {
				t.Errorf("Result differs. Actual: %v, Expected %v", actual, test.expected)
			}
		})
	}
}

func TestTouching(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		a        period.Period
		b        period.Period
		expected bool
	}{
		// nil cases
		{
			nil,
			nil,
			false,
		},
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			nil,
			false,
		},
		{
			nil,
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			false,
		},
		// Completely seperate
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			testPeriod{1, timestamp(0, 0, 3, 0), timestamp(0, 0, 4, 0)},
			false,
		},
		{
			testPeriod{1, timestamp(0, 0, 3, 0), timestamp(0, 0, 4, 0)},
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			false,
		},
		// Edges touching
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
			true,
		},
		{
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			true,
		},
		// Partial overlap
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 5, 0)},
			true,
		},
		{
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 5, 0)},
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
			true,
		},
		// One period contained in another
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 5, 0)},
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
			true,
		},
		{
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 5, 0)},
			true,
		},
		// Equal periods
		{
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
			true,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual := period.Touching(test.a, test.b)

			// Verify result
			if actual != test.expected {
				t.Errorf("Result differs. Actual: %v, Expected %v", actual, test.expected)
			}
		})
	}
}

func TestOverlap(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		a        period.Period
		b        period.Period
		expected time.Duration
	}{
		// nil cases
		{
			nil,
			nil,
			time.Duration(0),
		},
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			nil,
			time.Duration(0),
		},
		{
			nil,
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			time.Duration(0),
		},
		// Completely seperate
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			testPeriod{1, timestamp(0, 0, 3, 0), timestamp(0, 0, 4, 0)},
			time.Duration(0),
		},
		{
			testPeriod{1, timestamp(0, 0, 3, 0), timestamp(0, 0, 4, 0)},
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			time.Duration(0),
		},
		// Edges touching
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
			time.Duration(0),
		},
		{
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			time.Duration(0),
		},
		// Partial overlap
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 5, 0)},
			time.Duration(1 * time.Second),
		},
		{
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 5, 0)},
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
			time.Duration(1 * time.Second),
		},
		// One period contained in another
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 5, 0)},
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
			time.Duration(1 * time.Second),
		},
		{
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 5, 0)},
			time.Duration(1 * time.Second),
		},
		// Equal periods
		{
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
			time.Duration(1 * time.Second),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual := period.Overlap(test.a, test.b)

			// Verify result
			if actual != test.expected {
				t.Errorf("Result differs. Actual: %v, Expected %v", actual, test.expected)
			}
		})
	}
}

func TestShift(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		period   period.Period
		duration time.Duration
		expected period.Period
	}{
		// Nil cases
		{
			nil,
			0,
			nil,
		},
		{
			nil,
			1,
			nil,
		},
		// zero duration
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			0,
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
		},
		// Positive duration
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			time.Second,
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
		},
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			time.Minute,
			testPeriod{1, timestamp(0, 1, 1, 0), timestamp(0, 1, 2, 0)},
		},
		// Negative duration
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			-time.Second,
			testPeriod{1, timestamp(0, 0, 0, 0), timestamp(0, 0, 1, 0)},
		},
		{
			testPeriod{1, timestamp(0, 1, 1, 0), timestamp(0, 1, 2, 0)},
			-time.Minute,
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual := period.Shift(test.period, test.duration)

			// Verify result
			if actual != test.expected {
				t.Errorf("Result differs. Actual: %v, Expected %v", actual, test.expected)
			}
		})
	}
}

func TestScale(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		period   period.Period
		origin   time.Time
		factor   float64
		expected period.Period
	}{
		// Nil cases
		{
			nil,
			timestamp(0, 0, 0, 0),
			1.0,
			nil,
		},
		// Origin aligns with empty period
		{
			testPeriod{1, timestamp(0, 0, 0, 0), timestamp(0, 0, 0, 0)},
			timestamp(0, 0, 0, 0),
			0.0,
			testPeriod{1, timestamp(0, 0, 0, 0), timestamp(0, 0, 0, 0)},
		},
		{
			testPeriod{1, timestamp(0, 0, 0, 0), timestamp(0, 0, 0, 0)},
			timestamp(0, 0, 0, 0),
			1.0,
			testPeriod{1, timestamp(0, 0, 0, 0), timestamp(0, 0, 0, 0)},
		},
		{
			testPeriod{1, timestamp(0, 0, 0, 0), timestamp(0, 0, 0, 0)},
			timestamp(0, 0, 0, 0),
			2.0,
			testPeriod{1, timestamp(0, 0, 0, 0), timestamp(0, 0, 0, 0)},
		},
		{
			testPeriod{1, timestamp(0, 0, 0, 0), timestamp(0, 0, 0, 0)},
			timestamp(0, 0, 0, 0),
			-1.0,
			testPeriod{1, timestamp(0, 0, 0, 0), timestamp(0, 0, 0, 0)},
		},
		// Origin before empty period
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 1, 0)},
			timestamp(0, 0, 0, 0),
			0.0,
			testPeriod{1, timestamp(0, 0, 0, 0), timestamp(0, 0, 0, 0)},
		},
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 1, 0)},
			timestamp(0, 0, 0, 0),
			1.0,
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 1, 0)},
		},
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 1, 0)},
			timestamp(0, 0, 0, 0),
			2.0,
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 2, 0)},
		},
		{
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 2, 0)},
			timestamp(0, 0, 1, 0),
			-1.0,
			testPeriod{1, timestamp(0, 0, 0, 0), timestamp(0, 0, 0, 0)},
		},
		// Origin after empty period
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 1, 0)},
			timestamp(0, 0, 2, 0),
			0.0,
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 2, 0)},
		},
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 1, 0)},
			timestamp(0, 0, 2, 0),
			1.0,
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 1, 0)},
		},
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 1, 0)},
			timestamp(0, 0, 2, 0),
			2.0,
			testPeriod{1, timestamp(0, 0, 0, 0), timestamp(0, 0, 0, 0)},
		},
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 1, 0)},
			timestamp(0, 0, 2, 0),
			-1.0,
			testPeriod{1, timestamp(0, 0, 3, 0), timestamp(0, 0, 3, 0)},
		},
		// Origin before period
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			timestamp(0, 0, 0, 0),
			0.0,
			testPeriod{1, timestamp(0, 0, 0, 0), timestamp(0, 0, 0, 0)},
		},
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			timestamp(0, 0, 0, 0),
			1.0,
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
		},
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			timestamp(0, 0, 0, 0),
			2.0,
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
		},
		{
			testPeriod{1, timestamp(0, 0, 4, 0), timestamp(0, 0, 5, 0)},
			timestamp(0, 0, 3, 0),
			-1.0,
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
		},
		// Origin on period start
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			timestamp(0, 0, 1, 0),
			0.0,
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 1, 0)},
		},
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			timestamp(0, 0, 1, 0),
			1.0,
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
		},
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			timestamp(0, 0, 1, 0),
			2.0,
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
		},
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			timestamp(0, 0, 1, 0),
			-1.0,
			testPeriod{1, timestamp(0, 0, 0, 0), timestamp(0, 0, 1, 0)},
		},
		// Origin on period midpoint
		{
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
			timestamp(0, 0, 3, 0),
			0.0,
			testPeriod{1, timestamp(0, 0, 3, 0), timestamp(0, 0, 3, 0)},
		},
		{
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
			timestamp(0, 0, 3, 0),
			1.0,
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
		},
		{
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
			timestamp(0, 0, 3, 0),
			2.0,
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 5, 0)},
		},
		{
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
			timestamp(0, 0, 3, 0),
			-1.0,
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
		},
		// Origin on period end
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			timestamp(0, 0, 2, 0),
			0.0,
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 2, 0)},
		},
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			timestamp(0, 0, 2, 0),
			1.0,
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
		},
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			timestamp(0, 0, 2, 0),
			2.0,
			testPeriod{1, timestamp(0, 0, 0, 0), timestamp(0, 0, 2, 0)},
		},
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			timestamp(0, 0, 2, 0),
			-1.0,
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
		},
		// Origin after period
		{
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
			timestamp(0, 0, 4, 0),
			0.0,
			testPeriod{1, timestamp(0, 0, 4, 0), timestamp(0, 0, 4, 0)},
		},
		{
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
			timestamp(0, 0, 4, 0),
			1.0,
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
		},
		{
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
			timestamp(0, 0, 4, 0),
			2.0,
			testPeriod{1, timestamp(0, 0, 0, 0), timestamp(0, 0, 2, 0)},
		},
		{
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
			timestamp(0, 0, 4, 0),
			-1.0,
			testPeriod{1, timestamp(0, 0, 5, 0), timestamp(0, 0, 6, 0)},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual := period.Scale(test.period, test.origin, test.factor)

			// Verify result
			if actual != test.expected {
				t.Errorf("Result differs. Actual: %v, Expected %v", actual, test.expected)
			}
		})
	}
}

func TestMin(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		a        period.Period
		b        period.Period
		timeFunc period.TimeFunction
		expected time.Time
	}{
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			testPeriod{2, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
			period.Period.End,
			timestamp(0, 0, 2, 0),
		},
		{
			testPeriod{1, timestamp(0, 0, 3, 0), timestamp(0, 0, 1, 0)},
			testPeriod{2, timestamp(0, 0, 2, 0), timestamp(0, 0, 1, 0)},
			period.Period.Start,
			timestamp(0, 0, 2, 0),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual := period.Min(test.a, test.b, test.timeFunc)

			// Verify result
			if actual != test.expected {
				t.Errorf("Result differs. Actual: %v, Expected %v", actual, test.expected)
			}
		})
	}
}

func TestMax(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		a        period.Period
		b        period.Period
		timeFunc period.TimeFunction
		expected time.Time
	}{
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			testPeriod{2, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
			period.Period.End,
			timestamp(0, 0, 3, 0),
		},
		{
			testPeriod{1, timestamp(0, 0, 3, 0), timestamp(0, 0, 1, 0)},
			testPeriod{2, timestamp(0, 0, 2, 0), timestamp(0, 0, 1, 0)},
			period.Period.Start,
			timestamp(0, 0, 3, 0),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual := period.Max(test.a, test.b, test.timeFunc)

			// Verify result
			if actual != test.expected {
				t.Errorf("Result differs. Actual: %v, Expected %v", actual, test.expected)
			}
		})
	}
}

func TestDuration(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		period   period.Period
		expected time.Duration
	}{
		{
			testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			time.Second,
		},
		{
			testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 2, 0)},
			0,
		},
		{
			nil,
			0,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual := period.Duration(test.period)

			// Verify result
			if actual != test.expected {
				t.Errorf("Result differs. Actual: %v, Expected %v", actual, test.expected)
			}
		})
	}
}

func timestamp(hour, min, sec, milli int) time.Time {
	return time.Date(0, time.January, 1, hour, min, sec, milli*1e+6, time.UTC)
}

type testPeriod struct {
	id    int
	start time.Time
	end   time.Time
}

func (tp testPeriod) Valid() bool {
	return tp.start.Before(tp.end) || tp.start.Equal(tp.end)
}

func (tp testPeriod) Start() time.Time {
	return tp.start
}

func (tp testPeriod) End() time.Time {
	return tp.end
}

func (tp testPeriod) TransformToNew(start, end time.Time) period.Period {
	return testPeriod{
		id:    tp.id,
		start: start,
		end:   end,
	}
}
