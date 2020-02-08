package period

import (
	"time"

	cabiriaTime "github.com/liampulles/cabiria/pkg/time"
)

// TimeFunction must return a time value for a given Period.
//  Generally, this will be Period.Start or Period.End.
type TimeFunction func(Period) time.Time

// DoesOverlap returns true if a and b overlap, otherwise false.
//  a and b are NOT considered to be overlapping if their bounds merely
//  touch - for that see Touching.
func DoesOverlap(a, b Period) bool {
	if a == nil || b == nil {
		return false
	}
	return a.Start().Before(b.End()) && b.Start().Before(a.End())
}

// Touching returns true if a and b overlap OR if their bounds touch,
//  otherwise false.
func Touching(a, b Period) bool {
	if a == nil || b == nil {
		return false
	}
	return !b.End().Before(a.Start()) && !a.End().Before(b.Start())
}

// Overlap returns the duration of the section for which a and b overlap.
//  If a and b do NOT overlap, or either is nil, then 0 is returned.
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

// Shift returns a new period with the start and end adjusted to be +amount.
func Shift(period Period, amount time.Duration) Period {
	if period == nil {
		return nil
	}
	newStart := period.Start().Add(amount)
	newEnd := period.End().Add(amount)
	return period.TransformToNew(newStart, newEnd)
}

// Scale returns a new period where period has been scaled by factor from origin.
// e.g. period: (0:00:02.000,0:00:03.000), origin: 0:00:01.000, factor: 2.0 => (0:00:03.000,0:00:05.000)
func Scale(period Period, origin time.Time, factor float64) Period {
	if period == nil {
		return nil
	}
	newStart := cabiriaTime.Scale(period.Start(), origin, factor)
	newEnd := cabiriaTime.Scale(period.End(), origin, factor)
	// If the scale is negative, switch them.
	if factor < 0.0 {
		newStart, newEnd = newEnd, newStart
	}
	return period.TransformToNew(newStart, newEnd)
}

// Min returns the minimum time of timeFunc(a) vs. timeFunc(b)
func Min(a, b Period, timeFunc TimeFunction) time.Time {
	return cabiriaTime.Min(timeFunc(a), timeFunc(b))
}

// Max returns the maximum time of timeFunc(a) vs. timeFunc(b)
func Max(a, b Period, timeFunc TimeFunction) time.Time {
	return cabiriaTime.Max(timeFunc(a), timeFunc(b))
}

// Duration returns the duration that a period covers. If period is nil,
//  0 is returned.
func Duration(period Period) time.Duration {
	if period == nil {
		return 0
	}
	return period.End().Sub(period.Start())
}
