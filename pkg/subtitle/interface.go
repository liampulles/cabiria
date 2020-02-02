package subtitle

import (
	"time"
)

type Subtitle struct {
	Start time.Time
	End   time.Time
	Text  string
}
