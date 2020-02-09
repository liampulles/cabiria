package time_test

import (
	"fmt"
	"testing"
	"time"

	cabiriaTime "github.com/liampulles/cabiria/pkg/time"
)

func TestFromFrameAndFPS(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		frame    int
		fps      float64
		expected time.Time
	}{
		{
			0,
			0.0,
			time.Time{},
		},
		{
			0,
			1.0,
			timestamp(0, 0, 0, 0),
		},
		{
			1,
			0.0,
			time.Time{},
		},
		{
			1,
			1.0,
			timestamp(0, 0, 1, 0),
		},
		{
			1,
			2.0,
			timestamp(0, 0, 0, 500),
		},
		{
			2,
			2.0,
			timestamp(0, 0, 1, 0),
		},

		{
			24536,
			25.000,
			timestamp(0, 16, 21, 440),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual := cabiriaTime.FromFrameAndFPS(test.frame, test.fps)

			// Verify result
			if !actual.Equal(test.expected) {
				t.Errorf("Result differs. Actual: %s, Expected %s", actual, test.expected)
			}
		})
	}
}

func timestamp(hour, min, sec, milli int) time.Time {
	return time.Date(0, time.January, 1, hour, min, sec, milli*1e+6, time.UTC)
}
