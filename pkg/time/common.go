package time

import "time"

// Min returns the time that is earliest.
func Min(a, b time.Time) time.Time {
	if a.Before(b) {
		return a
	}
	return b
}

// Max returns the time that is latest.
func Max(a, b time.Time) time.Time {
	if a.Before(b) {
		return b
	}
	return a
}

// Scale scales t from origin by factor.
//  e.g. t: 0:00:01.000, origin: 0:00:02.000, factor: 2.0 -> 0:00:03.000
func Scale(t time.Time, origin time.Time, factor float64) time.Time {
	currentDist := t.Sub(origin)
	newDist := time.Duration(float64(currentDist) * factor)
	return origin.Add(newDist)
}
