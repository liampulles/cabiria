package read_test

import (
	"fmt"
	"path"
	"testing"
	"time"

	"github.com/liampulles/cabiria/pkg/subtitle"
	"github.com/liampulles/cabiria/pkg/subtitle/read"
	subTest "github.com/liampulles/cabiria/pkg/subtitle/test"
)

func TestReadSRT_WhenSRTValid(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		path     string
		expected []subtitle.Subtitle
	}{
		{
			"blank.1.srt",
			subs(),
		},
		{
			"blank.2.srt",
			subs(),
		},
		{
			"single.srt",
			subs(
				sub(
					"First line\nSecond line",
					timestamp(0, 2, 31, 567),
					timestamp(0, 2, 37, 164),
				),
			),
		},
		{
			"many.srt",
			subs(
				sub(
					"First line\nSecond line",
					timestamp(0, 1, 11, 111),
					timestamp(0, 2, 22, 222),
				),
				sub(
					"First line\nSecond line",
					timestamp(0, 2, 22, 222),
					timestamp(0, 3, 33, 333),
				),
				sub(
					"First line\nSecond line",
					timestamp(0, 0, 0, 0),
					timestamp(23, 59, 59, 999),
				),
			),
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("[%s]", test.path), func(t *testing.T) {
			// Exercise SUT
			actual, err := read.ReadSRT(path.Join("testdata", test.path))

			// Verify result
			if err != nil {
				t.Errorf("Encountered exception in SUT: %v", err)
			}
			if err = subTest.CompareSubtitles(actual, test.expected); err != nil {
				t.Errorf("Comparison failure: %v", err)
			}
		})
	}
}

func subs(subs ...subtitle.Subtitle) []subtitle.Subtitle {
	return subs
}

func sub(text string, start, end time.Time) subtitle.Subtitle {
	return subtitle.Subtitle{
		Start: start,
		End:   end,
		Text:  text,
	}
}

func timestamp(hour, min, sec, milli int) time.Time {
	return time.Date(0, time.January, 1, hour, min, sec, milli*1e+6, time.UTC)
}
