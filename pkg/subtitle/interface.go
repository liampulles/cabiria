package subtitle

import (
	"time"
)

type Style int

const (
	ItalicStyle Style = iota
	UnderlinedStyle
	BoldStyle
)

type Subtitle struct {
	Start time.Time
	End   time.Time
	Text  string
}
