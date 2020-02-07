package period

import "time"

type Period interface {
	Valid() bool
	Start() time.Time
	End() time.Time
	TransformToNew(time.Time, time.Time) Period
}
