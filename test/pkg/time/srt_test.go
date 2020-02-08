package time_test

import (
	"fmt"
	"testing"
	"time"

	cabiriaTime "github.com/liampulles/cabiria/pkg/time"
)

func TestFromSRTTimecode_WhenInputIsValid(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		timecode string
		expected time.Time
	}{
		{
			"00:00:00,000",
			timestamp(0, 0, 0, 0),
		},
		{
			"00:00:00,001",
			timestamp(0, 0, 0, 1),
		},
		{
			"00:00:01,000",
			timestamp(0, 0, 1, 0),
		},
		{
			"00:01:00,000",
			timestamp(0, 1, 0, 0),
		},
		{
			"01:00:00,000",
			timestamp(1, 0, 0, 0),
		},
		{
			"23:59:59,999",
			timestamp(23, 59, 59, 999),
		},
		{
			"12:34:56,789",
			timestamp(12, 34, 56, 789),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual, err := cabiriaTime.FromSRTTimecode(test.timecode)

			// Verify result
			if err != nil {
				t.Errorf("SUT threw an error: %v", err)
			}
			if !actual.Equal(test.expected) {
				t.Errorf("Result differs. Actual: %s, Expected %s", actual, test.expected)
			}
		})
	}
}
