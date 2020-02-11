package subtitle

import (
	"time"

	"github.com/liampulles/cabiria/pkg/time/period"
)

// Valid returns true if a subtitle is not a valid Period, otherwise false.
func (s Subtitle) Valid() bool {
	return !s.EndTime.Before(s.StartTime)
}

// Start returns the start of a subtitle
func (s Subtitle) Start() time.Time {
	return s.StartTime
}

// End returns the end of a subtitle
func (s Subtitle) End() time.Time {
	return s.EndTime
}

// TransformToNew returns a new subtitle which changes the start and end to
//  the desired times.
func (s Subtitle) TransformToNew(start, end time.Time) period.Period {
	return Subtitle{
		StartTime: start,
		EndTime:   end,
		Text:      s.Text,
		Style:     s.Style,
	}
}
