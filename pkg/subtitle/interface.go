package subtitle

import (
	"time"
)

type Subtitle struct {
	StartTime time.Time
	EndTime   time.Time
	Text      string
}
