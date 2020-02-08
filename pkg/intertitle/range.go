package intertitle

import (
	"time"

	cabiriaTime "github.com/liampulles/cabiria/pkg/time"
	"github.com/liampulles/cabiria/pkg/time/period"
)

// Range defines a set of frames which encapsulate an intertitle.
//  Range can be used as a Period.
type Range struct {
	StartFrame int
	EndFrame   int
	FPS        float64
}

// Valid will return true if a range is valid, otherwise false.
func (ir Range) Valid() bool {
	return ir.FPS > 0.0 &&
		ir.StartFrame >= 0 &&
		ir.EndFrame >= 0 &&
		ir.StartFrame <= ir.EndFrame
}

// Start returns a time representation of the start frame of a Range,
//  using the FPS.
func (ir Range) Start() time.Time {
	return cabiriaTime.FromFrameAndFPS(ir.StartFrame, ir.FPS)
}

// End returns a time representation of the end frame of a Range,
//  using the FPS.
func (ir Range) End() time.Time {
	return cabiriaTime.FromFrameAndFPS(ir.EndFrame, ir.FPS)
}

// TransformToNew computes a new Range given the desired start and end times,
//  calculating frame numbers using the FPS.
func (ir Range) TransformToNew(start, end time.Time) period.Period {
	return Range{
		StartFrame: fromTimeAndFPS(start, ir.FPS),
		EndFrame:   fromTimeAndFPS(end, ir.FPS),
		FPS:        ir.FPS,
	}
}

// MapRanges takes an array of intertitle frames and an fps, and reduces it
//  to an array of Ranges.
func MapRanges(intertitles []bool, fps float64) []Range {
	transitions := make([]Range, 0)
	last := false
	start := -1
	for i, current := range intertitles {
		// Start of intertitle
		if !last && current {
			start = i
		}
		// End of intertitle
		if last && !current {
			transitions = appendIntertitle(transitions, start, i-1, fps)
			start = -1
		}
		last = current
	}
	// Close off end, if applicable
	transitions = appendIntertitle(transitions, start, len(intertitles)-1, fps)
	return transitions
}

func appendIntertitle(transitions []Range, start, end int, fps float64) []Range {
	if start < 0 {
		return transitions
	}
	new := Range{
		StartFrame: start,
		EndFrame:   end,
		FPS:        fps,
	}
	return append(transitions, new)
}

func fromTimeAndFPS(t time.Time, fps float64) int {
	hours := time.Duration(t.Hour()) * time.Hour
	minutes := time.Duration(t.Minute()) * time.Minute
	seconds := time.Duration(t.Second()) * time.Second
	nano := time.Duration(t.Nanosecond()) * time.Nanosecond
	totalSeconds := float64(hours+minutes+seconds+nano) / float64(time.Second)
	return int(totalSeconds * fps)
}
