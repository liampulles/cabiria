package period

import (
	"time"

	calibriaTime "github.com/liampulles/cabiria/pkg/time"
)

type TimeFunction func(Period) time.Time

func DoesOverlap(a, b Period) bool {
	if a == nil || b == nil {
		return false
	}
	return a.Start().Before(b.End()) && b.Start().Before(a.End())
}

func Touching(a, b Period) bool {
	if a == nil || b == nil {
		return false
	}
	return !b.End().Before(a.Start()) && !a.End().Before(b.Start())
}

func Overlap(a, b Period) time.Duration {
	if a == nil || b == nil {
		return 0
	}
	latestStart := Max(a, b, Period.Start)
	earliestEnd := Min(a, b, Period.End)
	if earliestEnd.Before(latestStart) {
		return 0
	}
	return earliestEnd.Sub(latestStart)
}

func Shift(period Period, amount time.Duration) Period {
	if period == nil {
		return nil
	}
	newStart := period.Start().Add(amount)
	newEnd := period.End().Add(amount)
	return period.TransformToNew(newStart, newEnd)
}

func Scale(period Period, origin time.Time, factor float64) Period {
	if period == nil {
		return nil
	}
	newStart := calibriaTime.Scale(period.Start(), origin, factor)
	newEnd := calibriaTime.Scale(period.End(), origin, factor)
	// If the scale is negative, switch them.
	if factor < 0.0 {
		newStart, newEnd = newEnd, newStart
	}
	return period.TransformToNew(newStart, newEnd)
}

func Min(a, b Period, timeFunc TimeFunction) time.Time {
	return calibriaTime.Min(timeFunc(a), timeFunc(b))
}

func Max(a, b Period, timeFunc TimeFunction) time.Time {
	return calibriaTime.Max(timeFunc(a), timeFunc(b))
}

func Duration(period Period) time.Duration {
	if period == nil {
		return 0
	}
	return period.End().Sub(period.Start())
}
