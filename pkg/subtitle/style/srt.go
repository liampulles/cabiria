package style

import (
	"strings"
)

type styleTagPos struct {
	Start int
	End   int
}

// TODO: At the moment, encoding styles is more complexity then it is worth - so
// rather than storing the styles in SRT text, we just seek to remove any styling.
// One day...?
func RemoveStylesFromSRTText(text string) string {
	text2 := text
	lastLength := len(text2) + 1
	for lastLength != len(text2) {
		lastLength = len(text2)
		text2 = removeATag(text2)
	}
	return text2
}

func removeATag(text string) string {
	startPos := strings.Index(text, "<")
	if startPos < 0 {
		return text
	}
	endPos := strings.Index(text[startPos:], ">")
	if endPos < 0 {
		return text
	}
	return text[:startPos] + text[endPos+startPos+1:]
}
