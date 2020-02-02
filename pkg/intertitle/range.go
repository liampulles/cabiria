package intertitle

type IntertitleRange struct {
	StartFrame int
	EndFrame   int
}

func MapIntertitleRanges(intertitles []bool) []IntertitleRange {
	var transitions []IntertitleRange
	last := false
	start := -1
	for i, current := range intertitles {
		// Start of intertitle
		if !last && current {
			start = i
		}
		// End of intertitle
		if last && !current {
			appendIntertitle(transitions, start, i-1)
		}
		last = current
	}
	// If the end was a interitle, just close it off.
	if last {
		appendIntertitle(transitions, start, len(intertitles)-1)
	}
	return transitions
}

func appendIntertitle(transitions []IntertitleRange, start int, end int) {
	new := IntertitleRange{
		StartFrame: start,
		EndFrame:   end,
	}
	transitions = append(transitions, new)
}
