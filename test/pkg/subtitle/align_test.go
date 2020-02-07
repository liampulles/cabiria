package subtitle_test

import (
	"fmt"
	"testing"

	"github.com/liampulles/cabiria/pkg/intertitle"
	"github.com/liampulles/cabiria/pkg/subtitle"
	subTest "github.com/liampulles/cabiria/pkg/subtitle/test"
)

func TestAlignSubtitles(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		subs        []subtitle.Subtitle
		interRanges []intertitle.IntertitleRange
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
		// Already aigned intertitle and sub
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
		// Partially overlapping sub with touching intertitles
		{
			subs(
				sub(timestamp(0, 0, 1, 0), timestamp(0, 0, 3, 0), "text"),
			),
			interRanges(
				interRange(0, 2, 1.0),
				interRange(2, 4, 1.0),
			),
			subs(
				sub(timestamp(0, 0, 0, 0), timestamp(0, 0, 4, 0), "text"),
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
				sub(timestamp(0, 0, 1, 0), timestamp(0, 0, 5, 0), "text1"),
				sub(timestamp(0, 0, 6, 0), timestamp(0, 0, 10, 0), "text2"),
				sub(timestamp(0, 0, 11, 0), timestamp(0, 0, 15, 0), "text3"),
			),
			interRanges(
				interRange(0, 3, 1.0),
				interRange(4, 8, 1.0),
				interRange(9, 13, 1.0),
			),
			subs(
				sub(timestamp(0, 0, 0, 0), timestamp(0, 0, 3, 0), "text1"),
				sub(timestamp(0, 0, 4, 0), timestamp(0, 0, 8, 0), "text2"),
				sub(timestamp(0, 0, 9, 0), timestamp(0, 0, 13, 0), "text3"),
			),
		},
		// Offset case, where two subtitles fit into one intertitle
		{
			subs(
				sub(timestamp(0, 0, 1, 0), timestamp(0, 0, 5, 0), "text1"),
				sub(timestamp(0, 0, 6, 0), timestamp(0, 0, 10, 0), "text2"),
				sub(timestamp(0, 0, 11, 0), timestamp(0, 0, 15, 0), "text3"),
			),
			interRanges(
				interRange(0, 3, 1.0),
				interRange(4, 13, 1.0),
				interRange(14, 15, 1.0),
			),
			subs(
				sub(timestamp(0, 0, 0, 0), timestamp(0, 0, 3, 0), "text1"),
				sub(timestamp(0, 0, 4, 0), timestamp(0, 0, 8, 500), "text2"),
				sub(timestamp(0, 0, 8, 500), timestamp(0, 0, 13, 0), "text3"),
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
			fmt.Printf("%d", i)
			fmt.Printf("%d", i)
			// Exercise SUT
			actual := subtitle.AlignSubtitles(test.subs, test.interRanges)

			// Verify result
			if err := subTest.CompareSubtitles(actual, test.expected); err != nil {
				t.Errorf("Unexpected result: %v", err)
			}
		})
	}
}

func subs(subs ...subtitle.Subtitle) []subtitle.Subtitle {
	return subs
}

func interRange(start, end int, fps float64) intertitle.IntertitleRange {
	return intertitle.IntertitleRange{
		StartFrame: start,
		EndFrame:   end,
		FPS:        fps,
	}
}

func interRanges(irs ...intertitle.IntertitleRange) []intertitle.IntertitleRange {
	return irs
}
