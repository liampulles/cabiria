package intertitle_test

import (
	"fmt"
	"image/color"
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
				interRangeWithStyle(0, 0, 1.0, style(white(), white())),
			),
		},
		{
			intertitles(1, 0),
			1.0,
			interRanges(
				interRangeWithStyle(0, 0, 1.0, style(white(), white())),
			),
		},
		{
			intertitles(0, 1),
			1.0,
			interRanges(
				interRangeWithStyle(1, 1, 1.0, style(black(), black())),
			),
		},
		{
			intertitles(0, 1, 0),
			1.0,
			interRanges(
				interRangeWithStyle(1, 1, 1.0, style(black(), black())),
			),
		},
		// A multiple
		{
			intertitles(1, 1),
			1.0,
			interRanges(
				interRangeWithStyle(0, 1, 1.0, style(white(), white())),
			),
		},
		{
			intertitles(1, 1, 1),
			1.0,
			interRanges(
				interRangeWithStyle(0, 2, 1.0, style(black(), black())),
			),
		},
		{
			intertitles(1, 1, 0),
			1.0,
			interRanges(
				interRangeWithStyle(0, 1, 1.0, style(white(), white())),
			),
		},
		{
			intertitles(0, 1, 1),
			1.0,
			interRanges(
				interRangeWithStyle(1, 2, 1.0, style(black(), black())),
			),
		},
		{
			intertitles(0, 1, 1, 0),
			1.0,
			interRanges(
				interRangeWithStyle(1, 2, 1.0, style(black(), black())),
			),
		},
		// Complex case
		{
			intertitles(1, 0, 1, 1, 1, 0, 1, 0, 1, 1),
			1.0,
			interRanges(
				interRangeWithStyle(0, 0, 1.0, style(white(), white())),
				interRangeWithStyle(2, 4, 1.0, style(black(), black())),
				interRangeWithStyle(6, 6, 1.0, style(white(), white())),
				interRangeWithStyle(8, 9, 1.0, style(white(), white())),
			),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual, err := intertitle.MapRanges(test.intertitles, test.fps, framePaths())

			// Verify result
			if err != nil {
				t.Errorf("SUT returned an error: %v", err)
			}
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Result differs. Actual: %v, Expected %v", actual, test.expected)
			}
		})
	}
}

func framePaths() []string {
	var result []string
	for i := 0; i < 10; i++ {
		result = append(result, fmt.Sprintf("testdata/frames/%d.png", i))
	}
	return result
}

func interRange(start, end int, fps float64) intertitle.Range {
	return intertitle.Range{
		StartFrame: start,
		EndFrame:   end,
		FPS:        fps,
	}
}

func interRangeWithStyle(start, end int, fps float64, style intertitle.Style) intertitle.Range {
	return intertitle.Range{
		StartFrame: start,
		EndFrame:   end,
		FPS:        fps,
		Style:      style,
	}
}

func interRanges(interRanges ...intertitle.Range) []intertitle.Range {
	result := make([]intertitle.Range, 0)
	return append(result, interRanges...)
}

func timestamp(hour, min, sec, milli int) time.Time {
	return time.Date(0, time.January, 1, hour, min, sec, milli*1e+6, time.UTC)
}

func style(foreground, background color.Color) intertitle.Style {
	return intertitle.Style{
		ForegroundColor: foreground,
		BackgroundColor: background,
	}
}

func white() color.Color {
	return color.RGBA{
		255,
		255,
		255,
		255,
	}
}

func black() color.Color {
	return color.RGBA{
		0,
		0,
		0,
		255,
	}
}
