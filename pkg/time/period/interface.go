package period

import "time"

// Period is a conventional time period, which is a starting point
//  in time followed by an ending period of time, plus all the points in-between.
type Period interface {
	Valid() bool
	Start() time.Time
	End() time.Time
	TransformToNew(time.Time, time.Time) Period
}
