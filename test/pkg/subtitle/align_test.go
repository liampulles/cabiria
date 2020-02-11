package subtitle_test

import (
	"fmt"
	"image/color"
	"testing"
	"time"

	"github.com/liampulles/cabiria/pkg/intertitle"
	"github.com/liampulles/cabiria/pkg/subtitle"
	subTest "github.com/liampulles/cabiria/pkg/subtitle/test"
)

func TestAlignSubtitles(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		subs        []subtitle.Subtitle
		interRanges []intertitle.Range
		expected    []subtitle.Subtitle
	}{
		// No input, no output
		{
			subs(),
			interRanges(),
			subs(),
		},
		// No subs
		{
			subs(),
			interRanges(
				interRange(0, 1, 1.0),
			),
			subs(),
		},
		{
			subs(),
			interRanges(
				interRange(0, 1, 1.0),
				interRange(2, 5, 1.0),
			),
			subs(),
		},
		// No intertitles
		{
			subs(
				sub(timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0), "text"),
			),
			interRanges(),
			subs(
				sub(timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0), "text"),
			),
		},
		// -> Re-order subs
		{
			subs(
				sub(timestamp(0, 0, 3, 0), timestamp(0, 0, 4, 0), "text"),
				sub(timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0), "text"),
			),
			interRanges(),
			subs(
				sub(timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0), "text"),
				sub(timestamp(0, 0, 3, 0), timestamp(0, 0, 4, 0), "text"),
			),
		},
		// -> Fix overlapping subs
		{
			subs(
				sub(timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0), "text"),
				sub(timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0), "text"),
			),
			interRanges(),
			subs(
				sub(timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 500), "text"),
				sub(timestamp(0, 0, 2, 500), timestamp(0, 0, 4, 0), "text"),
			),
		},
		// Already aligned intertitle and sub
		{
			subs(
				sub(timestamp(0, 0, 0, 0), timestamp(0, 0, 1, 0), "text"),
			),
			interRanges(
				interRange(0, 1, 1.0),
			),
			subs(
				sub(timestamp(0, 0, 0, 0), timestamp(0, 0, 1, 0), "text"),
			),
		},
		// Not-at-all overlapping intertitle and sub
		{
			subs(
				sub(timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0), "text"),
			),
			interRanges(
				interRange(0, 1, 1.0),
			),
			subs(
				sub(timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0), "text"),
			),
		},
		{
			subs(
				sub(timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0), "text"),
			),
			interRanges(
				interRange(2, 3, 1.0),
			),
			subs(
				sub(timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0), "text"),
			),
		},
		// Partially overlapping intertitle and sub
		{
			subs(
				sub(timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0), "text"),
			),
			interRanges(
				interRange(0, 2, 1.0),
			),
			subs(
				sub(timestamp(0, 0, 0, 0), timestamp(0, 0, 2, 0), "text"),
			),
		},
		{
			subs(
				sub(timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0), "text"),
			),
			interRanges(
				interRange(2, 4, 1.0),
			),
			subs(
				sub(timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0), "text"),
			),
		},
		// Many touching subs overlapping intertitle
		{
			subs(
				sub(timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0), "text1"),
				sub(timestamp(0, 0, 2, 0), timestamp(0, 0, 3, 0), "text2"),
			),
			interRanges(
				interRange(0, 4, 1.0),
			),
			subs(
				sub(timestamp(0, 0, 0, 0), timestamp(0, 0, 2, 0), "text1"),
				sub(timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0), "text2"),
			),
		},
		// Many close subs overlapping intertitle
		{
			subs(
				sub(timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0), "text1"),
				sub(timestamp(0, 0, 3, 0), timestamp(0, 0, 4, 0), "text2"),
			),
			interRanges(
				interRange(0, 4, 1.0),
			),
			subs(
				sub(timestamp(0, 0, 0, 0), timestamp(0, 0, 2, 0), "text1"),
				sub(timestamp(0, 0, 2, 0), timestamp(0, 0, 4, 0), "text2"),
			),
		},
		// Offset subtitles case
		{
			subs(
				sub(timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0), "text1"),
				sub(timestamp(0, 0, 5, 0), timestamp(0, 0, 7, 0), "text2"),
				sub(timestamp(0, 0, 9, 0), timestamp(0, 0, 11, 0), "text3"),
			),
			interRanges(
				interRange(0, 2, 1.0),
				interRange(4, 6, 1.0),
				interRange(8, 10, 1.0),
			),
			subs(
				sub(timestamp(0, 0, 0, 0), timestamp(0, 0, 2, 0), "text1"),
				sub(timestamp(0, 0, 4, 0), timestamp(0, 0, 6, 0), "text2"),
				sub(timestamp(0, 0, 8, 0), timestamp(0, 0, 10, 0), "text3"),
			),
		},
		// Offset subtitles case, where they make an overlapping set
		{
			subs(
				sub(timestamp(0, 0, 1, 0), timestamp(0, 0, 6, 0), "text1"),
				sub(timestamp(0, 0, 7, 0), timestamp(0, 0, 11, 0), "text2"),
				sub(timestamp(0, 0, 12, 0), timestamp(0, 0, 15, 0), "text3"),
			),
			interRanges(
				interRange(0, 3, 1.0),
				interRange(5, 8, 1.0),
				interRange(10, 13, 1.0),
			),
			subs(
				sub(timestamp(0, 0, 0, 0), timestamp(0, 0, 3, 0), "text1"),
				sub(timestamp(0, 0, 5, 0), timestamp(0, 0, 8, 0), "text2"),
				sub(timestamp(0, 0, 10, 0), timestamp(0, 0, 13, 0), "text3"),
			),
		},
		// Offset case, where two subtitles fit into one intertitle
		{
			subs(
				sub(timestamp(0, 0, 1, 0), timestamp(0, 0, 6, 0), "text1"),
				sub(timestamp(0, 0, 7, 0), timestamp(0, 0, 10, 0), "text2"),
				sub(timestamp(0, 0, 11, 0), timestamp(0, 0, 14, 0), "text3"),
			),
			interRanges(
				interRange(0, 3, 1.0),
				interRange(5, 15, 1.0),
			),
			subs(
				sub(timestamp(0, 0, 0, 0), timestamp(0, 0, 3, 0), "text1"),
				sub(timestamp(0, 0, 5, 0), timestamp(0, 0, 10, 0), "text2"),
				sub(timestamp(0, 0, 10, 0), timestamp(0, 0, 15, 0), "text3"),
			),
		},
		// Offset case, where a subtitle is not overlapping an intertitle at all
		{
			subs(
				sub(timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0), "text1"),
				sub(timestamp(0, 0, 5, 0), timestamp(0, 0, 7, 0), "text2"),
				sub(timestamp(0, 0, 9, 0), timestamp(0, 0, 11, 0), "text3"),
				sub(timestamp(0, 0, 13, 0), timestamp(0, 0, 15, 0), "text4"),
			),
			interRanges(
				interRange(0, 2, 1.0),
				interRange(4, 6, 1.0),
				interRange(8, 10, 1.0),
			),
			subs(
				sub(timestamp(0, 0, 0, 0), timestamp(0, 0, 2, 0), "text1"),
				sub(timestamp(0, 0, 4, 0), timestamp(0, 0, 6, 0), "text2"),
				sub(timestamp(0, 0, 8, 0), timestamp(0, 0, 10, 0), "text3"),
				sub(timestamp(0, 0, 13, 0), timestamp(0, 0, 15, 0), "text4"),
			),
		},
		// Offset case, where an intertitle is not overlapping a subtitle at all
		{
			subs(
				sub(timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0), "text1"),
				sub(timestamp(0, 0, 5, 0), timestamp(0, 0, 7, 0), "text2"),
				sub(timestamp(0, 0, 9, 0), timestamp(0, 0, 11, 0), "text3"),
			),
			interRanges(
				interRange(0, 2, 1.0),
				interRange(4, 6, 1.0),
				interRange(8, 10, 1.0),
				interRange(12, 15, 1.0),
			),
			subs(
				sub(timestamp(0, 0, 0, 0), timestamp(0, 0, 2, 0), "text1"),
				sub(timestamp(0, 0, 4, 0), timestamp(0, 0, 6, 0), "text2"),
				sub(timestamp(0, 0, 8, 0), timestamp(0, 0, 10, 0), "text3"),
			),
		},
		// Offset case, where a subtitle overlaps another and is subsequently aligned
		{
			subs(
				sub(timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0), "text1"),
				sub(timestamp(0, 0, 5, 0), timestamp(0, 0, 7, 0), "text2"),
				sub(timestamp(0, 0, 9, 0), timestamp(0, 0, 12, 0), "text3"),
				sub(timestamp(0, 0, 11, 0), timestamp(0, 0, 14, 0), "text4"),
			),
			interRanges(
				interRange(0, 2, 1.0),
				interRange(4, 6, 1.0),
				interRange(8, 10, 1.0),
			),
			subs(
				sub(timestamp(0, 0, 0, 0), timestamp(0, 0, 2, 0), "text1"),
				sub(timestamp(0, 0, 4, 0), timestamp(0, 0, 6, 0), "text2"),
				sub(timestamp(0, 0, 8, 0), timestamp(0, 0, 9, 0), "text3"),
				sub(timestamp(0, 0, 9, 0), timestamp(0, 0, 10, 0), "text4"),
			),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual := subtitle.AlignSubtitles(test.subs, test.interRanges)

			// We don't want to test styles here, there is another test for that.
			actual = eraseStyles(actual)

			// Verify result
			if err := subTest.CompareSubtitles(actual, test.expected); err != nil {
				t.Errorf("Unexpected result: %v", err)
			}
		})
	}
}

func TestAlignSubtitles_ShouldCopyStyles(t *testing.T) {
	// Setup fixture
	inputSubs := subs(
		sub(timestamp(0, 0, 1, 0), timestamp(0, 0, 6, 0), "text1"),
		sub(timestamp(0, 0, 7, 0), timestamp(0, 0, 10, 0), "text2"),
		sub(timestamp(0, 0, 11, 0), timestamp(0, 0, 14, 0), "text3"),
	)
	inputInterRanges := interRanges(
		interRangeWithStyle(0, 3, 1.0, style(color.White, color.White)),
		interRangeWithStyle(5, 15, 1.0, style(color.Black, color.Black)),
	)

	// Setup expectations
	expectedSubs := subs(
		subWithStyle(timestamp(0, 0, 0, 0), timestamp(0, 0, 3, 0), "text1", style(color.White, color.White)),
		subWithStyle(timestamp(0, 0, 5, 0), timestamp(0, 0, 10, 0), "text2", style(color.Black, color.Black)),
		subWithStyle(timestamp(0, 0, 10, 0), timestamp(0, 0, 15, 0), "text3", style(color.Black, color.Black)),
	)

	// Exercise SUT
	actual := subtitle.AlignSubtitles(inputSubs, inputInterRanges)

	// Verify result
	if err := subTest.CompareSubtitles(actual, expectedSubs); err != nil {
		t.Errorf("Unexpected result: %v", err)
	}
}

func subs(subs ...subtitle.Subtitle) []subtitle.Subtitle {
	return subs
}

func subWithStyle(start, end time.Time, text string, style intertitle.Style) subtitle.Subtitle {
	return subtitle.Subtitle{
		StartTime: start,
		EndTime:   end,
		Text:      text,
		Style:     style,
	}
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

func interRanges(irs ...intertitle.Range) []intertitle.Range {
	result := make([]intertitle.Range, 0)
	return append(result, irs...)
}

func style(foreground, background color.Color) intertitle.Style {
	return intertitle.Style{
		ForegroundColor: foreground,
		BackgroundColor: background,
	}
}

func eraseStyles(subs []subtitle.Subtitle) []subtitle.Subtitle {
	if subs == nil {
		return nil
	}
	result := make([]subtitle.Subtitle, len(subs))
	for i, elem := range subs {
		elem.Style = intertitle.Style{}
		result[i] = elem
	}
	return result
}
