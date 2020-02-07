package time

import "time"

func Min(a, b time.Time) time.Time {
	if a.Before(b) {
		return a
	}
	return b
}

func Max(a, b time.Time) time.Time {
	if a.Before(b) {
		return b
	}
	return a
}

func Scale(t time.Time, origin time.Time, factor float64) time.Time {
	currentDist := t.Sub(origin)
	newDist := time.Duration(float64(currentDist) * factor)
	return origin.Add(newDist)
}
