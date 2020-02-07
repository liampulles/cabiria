package time

import (
	"time"
)

var InvalidTime = time.Date(-9999, -9999, -9999, -9999, -9999, -9999, -9999, time.UTC)

func FromFrameAndFPS(frame int, fps float64) time.Time {
	if fps == 0.0 {
		return InvalidTime
	}
	secDec := float64(frame) / fps
	whole := int64(secDec)
	nano := int64((secDec - float64(whole)) * 1e+9)
	base := time.Date(0, time.January, 1, 0, 0, 0, 0, time.UTC)
	base = base.Add(time.Duration(whole) * time.Second)
	base = base.Add(time.Duration(nano) * time.Nanosecond)
	return base
}
