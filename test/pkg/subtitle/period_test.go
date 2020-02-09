package subtitle_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/liampulles/cabiria/pkg/subtitle"
)

func TestValid(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		subtitle subtitle.Subtitle
		expected bool
	}{
		// Invalid cases
		{
			sub(timestamp(0, 0, 1, 0), timestamp(0, 0, 0, 0), "text"),
			false,
		},
		// Valid cases
		{
			sub(timestamp(0, 0, 0, 0), timestamp(0, 0, 0, 0), "text"),
			true,
		},
		{
			sub(timestamp(0, 0, 0, 0), timestamp(0, 0, 2, 0), "text"),
			true,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual := test.subtitle.Valid()

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
		subtitle subtitle.Subtitle
		expected time.Time
	}{
		{
			sub(timestamp(0, 0, 0, 0), timestamp(0, 0, 0, 0), "text"),
			timestamp(0, 0, 0, 0),
		},
		{
			sub(timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0), "text"),
			timestamp(0, 0, 1, 0),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual := test.subtitle.Start()

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
		subtitle subtitle.Subtitle
		expected time.Time
	}{
		{
			sub(timestamp(0, 0, 0, 0), timestamp(0, 0, 0, 0), "text"),
			timestamp(0, 0, 0, 0),
		},
		{
			sub(timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0), "text"),
			timestamp(0, 0, 2, 0),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual := test.subtitle.End()

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
		subtitle subtitle.Subtitle
		start    time.Time
		end      time.Time
		expected subtitle.Subtitle
	}{
		{
			sub(timestamp(0, 0, 0, 0), timestamp(0, 0, 0, 0), "text"),
			timestamp(0, 0, 0, 0),
			timestamp(0, 0, 0, 0),
			sub(timestamp(0, 0, 0, 0), timestamp(0, 0, 0, 0), "text"),
		},
		{
			sub(timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0), "text"),
			timestamp(0, 0, 4, 0),
			timestamp(0, 0, 6, 0),
			sub(timestamp(0, 0, 4, 0), timestamp(0, 0, 6, 0), "text"),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual := test.subtitle.TransformToNew(test.start, test.end)

			// Verify result
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Result differs. Actual: %v, Expected %v", actual, test.expected)
			}
		})
	}
}

func sub(start, end time.Time, text string) subtitle.Subtitle {
	return subtitle.Subtitle{
		StartTime: start,
		EndTime:   end,
		Text:      text,
	}
}

func timestamp(hour, min, sec, milli int) time.Time {
	return time.Date(0, time.January, 1, hour, min, sec, milli*1e+6, time.UTC)
}
