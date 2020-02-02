package subtitle

import (
	"time"

	"github.com/liampulles/cabiria/pkg/intertitle"
)

func AlignSubtitles(subs []Subtitle, interRanges []intertitle.IntertitleRange, fps float64) {
	for _, sub := range subs {
		overlapping := overlappingRanges(sub, interRanges, fps)
		alignSubtitle(sub, overlapping, fps)
	}
}

func overlappingRanges(sub Subtitle, interRanges []intertitle.IntertitleRange, fps float64) []intertitle.IntertitleRange {
	var overlapping []intertitle.IntertitleRange
	for _, interRange := range interRanges {
		if doesOverlap(sub, interRange, fps) {
			overlapping = append(overlapping, interRange)
		}
	}
	return overlapping
}

func alignSubtitle(sub Subtitle, overlapping []intertitle.IntertitleRange, fps float64) {
	if len(overlapping) == 0 {
		return
	}

	startBasis := overlapping[0]
	sub.Start = asDuration(startBasis.StartFrame, fps)

	endBasis := overlapping[len(overlapping)-1]
	sub.End = asDuration(endBasis.EndFrame, fps)
}

func doesOverlap(sub Subtitle, interRange intertitle.IntertitleRange, fps float64) bool {
	return sub.Start.Before(asDuration(interRange.EndFrame, fps)) &&
		asDuration(interRange.StartFrame, fps).Before(sub.End)
}

func asDuration(frame int, fps float64) time.Time {
	secDec := float64(frame) * fps
	whole := int64(secDec)
	nano := int64((secDec - float64(whole)) * 1e+9)
	return time.Unix(whole, nano)
}
