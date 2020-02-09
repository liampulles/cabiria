package time

import (
	"fmt"
	"time"
)

// ToASSTimecode formats a time as a timecode which is appropriate
//  for use in an ASS file.
func ToASSTimecode(t time.Time) string {
	return fmt.Sprintf("%d", t.Hour()) +
		t.Format(":04:05.") +
		fmt.Sprintf("%03d", time.Duration(t.Nanosecond())/time.Millisecond)[:2]
}
