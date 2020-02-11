package subtitle

import (
	"time"

	"github.com/liampulles/cabiria/pkg/time/period"

	"github.com/liampulles/cabiria/pkg/intertitle"
)

// AlignSubtitles tries to align the given subtitles to the detected intertitles
//  such that when the subtitles are played back, they align exactly with each
//  intertitle segment in the film - barring some edge cases arising due to
//  imperfect data.
func AlignSubtitles(subs []Subtitle, interRanges []intertitle.Range) ([]Subtitle, []intertitle.Range) {
	// TODO: Regularize input here
	mergedInterRanges := intertitle.JoinTouchingRanges(interRanges)

	joined := rangedSortedSet(subs, interRangesAsPeriods(mergedInterRanges))
	overlaps := overlappingSets(joined)
	return alignSubtitlesFromOverlappingSets(overlaps), mergedInterRanges
}

func overlappingSets(sortedRanges []period.Period) [][]period.Period {
	if len(sortedRanges) == 0 {
		return nil
	}
	var overlappingSets [][]period.Period
	currentSet := []period.Period{sortedRanges[0]}
	for i := 1; i < len(sortedRanges); i++ {
		elem := sortedRanges[i]
		if period.DoesOverlap(elem, period.Periods(currentSet)) {
			// Add to current set
			currentSet = append(currentSet, elem)
		} else {
			// Close current set, and init new
			overlappingSets = append(overlappingSets, currentSet)
			currentSet = []period.Period{elem}
		}
	}
	// Close final set
	overlappingSets = append(overlappingSets, currentSet)

	return overlappingSets
}

func alignSubtitlesFromOverlappingSets(sets [][]period.Period) []Subtitle {
	var subs []Subtitle
	for _, elem := range sets {
		elemSubs := alignSubtitlesFromOverlappingSet(elem)
		subs = append(subs, elemSubs...)
	}
	return subs
}

func alignSubtitlesFromOverlappingSet(set []period.Period) []Subtitle {
	// separate into intertitleRange and subtitle sets
	var interRangePeriods period.Periods
	var subtitlePeriods period.Periods
	for _, elem := range set {
		switch v := elem.(type) {
		case intertitle.Range:
			interRangePeriods = append(interRangePeriods, v)
		case Subtitle:
			subtitlePeriods = append(subtitlePeriods, v)
		}
	}
	// -> If no intertitles, or no subs -> Fix and return subs. //TODO: Maybe nil?
	if len(interRangePeriods) == 0 || len(subtitlePeriods) == 0 {
		return periodsAsSubs(period.FixOverlaps(subtitlePeriods))
	}

	// Scale the subtitle set to match the intertitleRange set bounds
	subtitlePeriods = subtitlePeriods.TransformToNew(interRangePeriods.Start(), interRangePeriods.End()).(period.Periods)

	// For each sub, determine which intertitle they MOST overlap with,
	//  and add them to the "bucket" for that intertitle.
	overlapBuckets := make([][]period.Period, len(interRangePeriods))
	for _, sub := range subtitlePeriods {
		maxOverlap := time.Duration(-1)
		var maxIdx int
		for i, interRange := range interRangePeriods {
			if overlap := period.Overlap(sub, interRange); overlap > maxOverlap {
				maxOverlap = overlap
				maxIdx = i
			}
		}
		overlapBuckets[maxIdx] = append(overlapBuckets[maxIdx], sub)
	}

	// For each intertitle bucket,
	// -> Shift and scale subs in bucket to match intertitle bounds
	// -> Scale subs to cover gaps in their range
	// -> Add subs to the final set
	var result []period.Period
	for i, bucket := range overlapBuckets {
		if len(bucket) == 0 {
			continue
		}
		newStart := interRangePeriods[i].Start()
		newEnd := interRangePeriods[i].End()
		newSubs := period.Periods(bucket).TransformToNew(newStart, newEnd).(period.Periods)
		newSubs = period.CoverGaps(newSubs)
		result = append(result, newSubs...)
	}

	// Fix subs to not overlap
	result = period.FixOverlaps(result)

	// return final set
	return periodsAsSubs(result)
}

func rangedSortedSet(subs []Subtitle, interRangePeriods []period.Period) []period.Period {
	var rangedSet []period.Period

	rangedSet = append(rangedSet, subsAsPeriods(subs)...)
	rangedSet = append(rangedSet, interRangePeriods...)

	period.Sort(rangedSet)
	return rangedSet
}

func subsAsPeriods(subs []Subtitle) period.Periods {
	var result []period.Period
	for _, elem := range subs {
		result = append(result, elem)
	}
	return result
}

func periodsAsSubs(periods []period.Period) []Subtitle {
	var result []Subtitle
	for _, elem := range periods {
		result = append(result, elem.(Subtitle))
	}
	return result
}

func interRangesAsPeriods(interRanges []intertitle.Range) period.Periods {
	var result []period.Period
	for _, elem := range interRanges {
		result = append(result, elem)
	}
	return result
}

func periodsAsInterRanges(periods []period.Period) []intertitle.Range {
	var result []intertitle.Range
	for _, elem := range periods {
		result = append(result, elem.(intertitle.Range))
	}
	return result
}
