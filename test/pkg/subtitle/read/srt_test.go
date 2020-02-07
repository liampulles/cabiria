package read_test

import (
	"fmt"
	"path"
	"testing"
	"time"

	"github.com/liampulles/cabiria/pkg/subtitle"
	"github.com/liampulles/cabiria/pkg/subtitle/read"
	subTest "github.com/liampulles/cabiria/pkg/subtitle/test"
	calibriaTime "github.com/liampulles/cabiria/pkg/time"
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
					timestamp("00:02:31,567"),
					timestamp("00:02:37,164"),
				),
			),
		},
		{
			"many.srt",
			subs(
				sub(
					"First line\nSecond line",
					timestamp("00:01:11,111"),
					timestamp("00:02:22,222"),
				),
				sub(
					"First line\nSecond line",
					timestamp("00:02:22,222"),
					timestamp("00:03:33,333"),
				),
				sub(
					"First line\nSecond line",
					timestamp("00:00:00,000"),
					timestamp("23:59:59,999"),
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
		StartTime: start,
		EndTime:   end,
		Text:      text,
	}
}

func timestamp(s string) time.Time {
	t, err := calibriaTime.FromSRTTimecode(s)
	if err != nil {
		panic(err)
	}
	return t
}
