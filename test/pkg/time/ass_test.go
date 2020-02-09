package time_test

import (
	"fmt"
	"testing"
	"time"

	cabiriaTime "github.com/liampulles/cabiria/pkg/time"
)

func TestToASSTimecode(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		t        time.Time
		expected string
	}{
		{
			timestamp(0, 0, 0, 0),
			"0:00:00.00",
		},
		{
			timestamp(1, 23, 45, 678),
			"1:23:45.67",
		},
		{
			timestamp(12, 34, 56, 90),
			"12:34:56.09",
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s -> %s", test.t.String(), test.expected), func(t *testing.T) {
			// Exercise SUT
			actual := cabiriaTime.ToASSTimecode(test.t)

			// Verify result
			if actual != test.expected {
				t.Errorf("Result differs. Actual: %s, Expected %s", actual, test.expected)
			}
		})
	}
}
