package periods_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/liampulles/cabiria/pkg/time/period"
)

func TestValid(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		periods  period.Periods
		expected bool
	}{
		{
			nil,
			false,
		},
		{
			periods(),
			false,
		},
		{
			periods(
				testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 1, 0)},
			),
			false,
		},
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 1, 0)},
			),
			true,
		},
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			),
			true,
		},
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
				testPeriod{1, timestamp(0, 0, 4, 0), timestamp(0, 0, 3, 0)},
			),
			false,
		},
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
				testPeriod{1, timestamp(0, 0, 3, 0), timestamp(0, 0, 4, 0)},
			),
			true,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual := test.periods.Valid()

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
		periods  period.Periods
		expected time.Time
	}{
		{
			nil,
			time.Time{},
		},
		{
			periods(),
			time.Time{},
		},
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			),
			timestamp(0, 0, 1, 0),
		},
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
				testPeriod{1, timestamp(0, 0, 3, 0), timestamp(0, 0, 4, 0)},
			),
			timestamp(0, 0, 1, 0),
		},
		{
			periods(
				testPeriod{1, timestamp(0, 0, 3, 0), timestamp(0, 0, 4, 0)},
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			),
			timestamp(0, 0, 1, 0),
		},
		{
			periods(
				testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 7, 0)},
				testPeriod{1, timestamp(0, 0, 3, 0), timestamp(0, 0, 4, 0)},
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
				testPeriod{1, timestamp(0, 0, 4, 0), timestamp(0, 0, 5, 0)},
			),
			timestamp(0, 0, 1, 0),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual := test.periods.Start()

			// Verify result
			if actual != test.expected {
				t.Errorf("Result differs. Actual: %v, Expected %v", actual, test.expected)
			}
		})
	}
}

func TestEnd(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		periods  period.Periods
		expected time.Time
	}{
		{
			nil,
			time.Time{},
		},
		{
			periods(),
			time.Time{},
		},
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			),
			timestamp(0, 0, 2, 0),
		},
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
				testPeriod{1, timestamp(0, 0, 3, 0), timestamp(0, 0, 4, 0)},
			),
			timestamp(0, 0, 4, 0),
		},
		{
			periods(
				testPeriod{1, timestamp(0, 0, 3, 0), timestamp(0, 0, 4, 0)},
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			),
			timestamp(0, 0, 4, 0),
		},
		{
			periods(
				testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 2, 0)},
				testPeriod{1, timestamp(0, 0, 3, 0), timestamp(0, 0, 4, 0)},
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 7, 0)},
				testPeriod{1, timestamp(0, 0, 4, 0), timestamp(0, 0, 5, 0)},
			),
			timestamp(0, 0, 7, 0),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual := test.periods.End()

			// Verify result
			if actual != test.expected {
				t.Errorf("Result differs. Actual: %v, Expected %v", actual, test.expected)
			}
		})
	}
}

func TestTransformToNew(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		periods  period.Periods
		start    time.Time
		end      time.Time
		expected period.Periods
	}{
		// Empty cases
		{
			nil,
			timestamp(0, 0, 1, 0),
			timestamp(0, 0, 2, 0),
			nil,
		},
		{
			periods(),
			timestamp(0, 0, 1, 0),
			timestamp(0, 0, 2, 0),
			periods(),
		},
		// Single period
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			),
			timestamp(0, 0, 1, 0),
			timestamp(0, 0, 2, 0),
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			),
		},
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			),
			timestamp(0, 0, 0, 0),
			timestamp(0, 0, 1, 0),
			periods(
				testPeriod{1, timestamp(0, 0, 0, 0), timestamp(0, 0, 1, 0)},
			),
		},
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			),
			timestamp(0, 0, 2, 0),
			timestamp(0, 0, 3, 0),
			periods(
				testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
			),
		},
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			),
			timestamp(0, 0, 4, 0),
			timestamp(0, 0, 5, 0),
			periods(
				testPeriod{1, timestamp(0, 0, 4, 0), timestamp(0, 0, 5, 0)},
			),
		},
		// Many periods
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
				testPeriod{2, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
				testPeriod{3, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			),
			timestamp(0, 0, 1, 0),
			timestamp(0, 0, 2, 0),
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
				testPeriod{2, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
				testPeriod{3, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			),
		},
		{
			periods(
				testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
				testPeriod{2, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
				testPeriod{3, timestamp(0, 0, 3, 0), timestamp(0, 0, 5, 0)},
			),
			timestamp(0, 0, 1, 0),
			timestamp(0, 0, 5, 0),
			periods(
				testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
				testPeriod{2, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
				testPeriod{3, timestamp(0, 0, 3, 0), timestamp(0, 0, 5, 0)},
			),
		},
		{
			periods(
				testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
				testPeriod{2, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
				testPeriod{3, timestamp(0, 0, 3, 0), timestamp(0, 0, 5, 0)},
			),
			timestamp(0, 0, 5, 0),
			timestamp(0, 0, 9, 0),
			periods(
				testPeriod{1, timestamp(0, 0, 6, 0), timestamp(0, 0, 8, 0)},
				testPeriod{2, timestamp(0, 0, 5, 0), timestamp(0, 0, 7, 0)},
				testPeriod{3, timestamp(0, 0, 7, 0), timestamp(0, 0, 9, 0)},
			),
		},
		{
			periods(
				testPeriod{1, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
				testPeriod{2, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
				testPeriod{3, timestamp(0, 0, 3, 0), timestamp(0, 0, 5, 0)},
			),
			timestamp(0, 0, 5, 0),
			timestamp(0, 0, 13, 0),
			periods(
				testPeriod{1, timestamp(0, 0, 7, 0), timestamp(0, 0, 11, 0)},
				testPeriod{2, timestamp(0, 0, 5, 0), timestamp(0, 0, 9, 0)},
				testPeriod{3, timestamp(0, 0, 9, 0), timestamp(0, 0, 13, 0)},
			),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual := test.periods.TransformToNew(test.start, test.end)

			// Verify result
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Result differs. Actual: %v, Expected %v", actual, test.expected)
			}
		})
	}
}

func TestFixOverlaps(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		periods  period.Periods
		expected period.Periods
	}{
		// Empty cases
		{
			nil,
			nil,
		},
		{
			periods(),
			periods(),
		},
		// Cases without an overlap
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			),
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			),
		},
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
				testPeriod{2, timestamp(0, 0, 3, 0), timestamp(0, 0, 4, 0)},
			),
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
				testPeriod{2, timestamp(0, 0, 3, 0), timestamp(0, 0, 4, 0)},
			),
		},
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
				testPeriod{2, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
			),
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
				testPeriod{2, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
			),
		},
		// One overlap
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
				testPeriod{2, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
			),
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 500)},
				testPeriod{2, timestamp(0, 0, 2, 500), timestamp(0, 0, 4, 0)},
			),
		},
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
				testPeriod{2, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
			),
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
				testPeriod{2, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
			),
		},
		{
			periods(
				testPeriod{3, timestamp(0, 0, 3, 0), timestamp(0, 0, 5, 0)},
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
				testPeriod{2, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
			),
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
				testPeriod{2, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 500)},
				testPeriod{3, timestamp(0, 0, 3, 500), timestamp(0, 0, 5, 0)},
			),
		},
		// Many overlaps
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
				testPeriod{2, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
				testPeriod{3, timestamp(0, 0, 3, 0), timestamp(0, 0, 5, 0)},
				testPeriod{4, timestamp(0, 0, 5, 0), timestamp(0, 0, 7, 0)},
				testPeriod{5, timestamp(0, 0, 6, 0), timestamp(0, 0, 8, 0)},
			),
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
				testPeriod{2, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 500)},
				testPeriod{3, timestamp(0, 0, 3, 500), timestamp(0, 0, 5, 0)},
				testPeriod{4, timestamp(0, 0, 5, 0), timestamp(0, 0, 6, 500)},
				testPeriod{5, timestamp(0, 0, 6, 500), timestamp(0, 0, 8, 0)},
			),
		},
		{
			periods(
				testPeriod{4, timestamp(0, 0, 4, 0), timestamp(0, 0, 6, 0)},
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
				testPeriod{2, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
				testPeriod{3, timestamp(0, 0, 3, 0), timestamp(0, 0, 5, 0)},
			),
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 250)},
				testPeriod{2, timestamp(0, 0, 2, 250), timestamp(0, 0, 3, 500)},
				testPeriod{3, timestamp(0, 0, 3, 500), timestamp(0, 0, 4, 750)},
				testPeriod{4, timestamp(0, 0, 4, 750), timestamp(0, 0, 6, 0)},
			),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			fmt.Printf("%d", i)
			// Exercise SUT
			actual := period.FixOverlaps(test.periods)

			// Verify result
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Result differs. Actual: %v, Expected %v", actual, test.expected)
			}
		})
	}
}

func TestMergeTouching(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		periods  period.Periods
		expected period.Periods
	}{
		// Empty cases
		{
			nil,
			nil,
		},
		{
			periods(),
			periods(),
		},
		// Cases without a touch
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			),
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
			),
		},
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
				testPeriod{2, timestamp(0, 0, 3, 0), timestamp(0, 0, 4, 0)},
			),
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
				testPeriod{2, timestamp(0, 0, 3, 0), timestamp(0, 0, 4, 0)},
			),
		},
		// One touch
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
				testPeriod{2, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
			),
			periods(
				testPeriod{2, timestamp(0, 0, 1, 0), timestamp(0, 0, 4, 0)},
			),
		},
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
				testPeriod{2, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
			),
			periods(
				testPeriod{2, timestamp(0, 0, 1, 0), timestamp(0, 0, 4, 0)},
			),
		},
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
				testPeriod{2, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
			),
			periods(
				testPeriod{2, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
			),
		},
		{
			periods(
				testPeriod{3, timestamp(0, 0, 3, 0), timestamp(0, 0, 5, 0)},
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 1, 500)},
				testPeriod{2, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
			),
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 1, 500)},
				testPeriod{6, timestamp(0, 0, 2, 0), timestamp(0, 0, 5, 0)},
			),
		},
		// Many overlaps
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
				testPeriod{2, timestamp(0, 0, 2, 500), timestamp(0, 0, 4, 0)},
				testPeriod{3, timestamp(0, 0, 3, 0), timestamp(0, 0, 5, 0)},
				testPeriod{4, timestamp(0, 0, 5, 500), timestamp(0, 0, 7, 0)},
				testPeriod{5, timestamp(0, 0, 6, 0), timestamp(0, 0, 8, 0)},
			),
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
				testPeriod{6, timestamp(0, 0, 2, 500), timestamp(0, 0, 5, 0)},
				testPeriod{20, timestamp(0, 0, 5, 500), timestamp(0, 0, 8, 0)},
			),
		},
		{
			periods(
				testPeriod{4, timestamp(0, 0, 4, 0), timestamp(0, 0, 6, 0)},
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
				testPeriod{2, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
				testPeriod{3, timestamp(0, 0, 3, 0), timestamp(0, 0, 5, 0)},
			),
			periods(
				testPeriod{24, timestamp(0, 0, 1, 0), timestamp(0, 0, 6, 0)},
			),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			fmt.Printf("%d", i)
			// Exercise SUT
			actual := period.MergeTouching(test.periods, func(a, b period.Period) period.Period {
				aTestPeriod := a.(testPeriod)
				bTestPeriod := b.(testPeriod)
				return testPeriod{
					aTestPeriod.id * bTestPeriod.id,
					aTestPeriod.start,
					bTestPeriod.end,
				}
			})

			// Verify result
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Result differs. Actual: %v, Expected %v", actual, test.expected)
			}
		})
	}
}

func TestCoverGaps(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		periods  period.Periods
		expected period.Periods
	}{
		// Empty cases
		{
			nil,
			periods(),
		},
		{
			periods(),
			periods(),
		},
		// Single element
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
			),
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
			),
		},
		// Many elements, which overlap
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
				testPeriod{2, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
			),
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
				testPeriod{2, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
			),
		},
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 5, 0)},
				testPeriod{2, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
			),
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 5, 0)},
				testPeriod{2, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
			),
		},
		{
			periods(
				testPeriod{2, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
			),
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
				testPeriod{2, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
			),
		},
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
				testPeriod{2, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
				testPeriod{3, timestamp(0, 0, 3, 0), timestamp(0, 0, 5, 0)},
			),
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
				testPeriod{2, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
				testPeriod{3, timestamp(0, 0, 3, 0), timestamp(0, 0, 5, 0)},
			),
		},
		// Many elements, which touch
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
				testPeriod{2, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
			),
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
				testPeriod{2, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
			),
		},
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
				testPeriod{3, timestamp(0, 0, 3, 0), timestamp(0, 0, 4, 0)},
				testPeriod{2, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
			),
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
				testPeriod{2, timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0)},
				testPeriod{3, timestamp(0, 0, 3, 0), timestamp(0, 0, 4, 0)},
			),
		},
		// Many elements, which have gaps
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
				testPeriod{2, timestamp(0, 0, 3, 0), timestamp(0, 0, 4, 0)},
			),
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 500)},
				testPeriod{2, timestamp(0, 0, 2, 500), timestamp(0, 0, 4, 0)},
			),
		},
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
				testPeriod{2, timestamp(0, 0, 5, 0), timestamp(0, 0, 7, 0)},
			),
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
				testPeriod{2, timestamp(0, 0, 3, 0), timestamp(0, 0, 7, 0)},
			),
		},
		{
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0)},
				testPeriod{2, timestamp(0, 0, 4, 0), timestamp(0, 0, 5, 0)},
				testPeriod{3, timestamp(0, 0, 7, 0), timestamp(0, 0, 9, 0)},
			),
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
				testPeriod{2, timestamp(0, 0, 3, 0), timestamp(0, 0, 6, 0)},
				testPeriod{3, timestamp(0, 0, 6, 0), timestamp(0, 0, 9, 0)},
			),
		},
		// Mixed case
		{
			periods(
				testPeriod{5, timestamp(0, 0, 9, 0), timestamp(0, 0, 11, 0)},
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
				testPeriod{4, timestamp(0, 0, 8, 0), timestamp(0, 0, 10, 0)},
				testPeriod{2, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
				testPeriod{3, timestamp(0, 0, 4, 0), timestamp(0, 0, 6, 0)},
			),
			periods(
				testPeriod{1, timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0)},
				testPeriod{2, timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0)},
				testPeriod{3, timestamp(0, 0, 4, 0), timestamp(0, 0, 7, 0)},
				testPeriod{4, timestamp(0, 0, 7, 0), timestamp(0, 0, 10, 0)},
				testPeriod{5, timestamp(0, 0, 9, 0), timestamp(0, 0, 11, 0)},
			),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			fmt.Printf("%d", i)
			// Exercise SUT
			actual := period.CoverGaps(test.periods)

			// Verify result
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Result differs. Actual: %v, Expected %v", actual, test.expected)
			}
		})
	}
}

func periods(periods ...period.Period) period.Periods {
	result := make(period.Periods, 0)
	return append(result, periods...)
}
