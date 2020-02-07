package subtitle

import (
	"time"

	"github.com/liampulles/cabiria/pkg/time/period"
)

func (s Subtitle) Valid() bool {
	return !s.EndTime.Before(s.StartTime)
}

func (s Subtitle) Start() time.Time {
	return s.StartTime
}

func (s Subtitle) End() time.Time {
	return s.EndTime
}

func (s Subtitle) TransformToNew(start, end time.Time) period.Period {
	return Subtitle{
		StartTime: start,
		EndTime:   end,
		Text:      s.Text,
	}
}
