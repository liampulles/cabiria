package subtitle

import (
	"time"

	"github.com/liampulles/cabiria/pkg/intertitle"
)

// Subtitle defines a single subtitle in a set of subtitles, that
// being when it starts and ends, and what text it displays.
type Subtitle struct {
	StartTime time.Time
	EndTime   time.Time
	Text      string
	Style     intertitle.Style
}
