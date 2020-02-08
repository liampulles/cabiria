package period

import (
	"sort"
	"time"

	cabiriaTime "github.com/liampulles/cabiria/pkg/time"
)

// A group of periods can itself be considered a period - we'll implement the
// interface here to demonstrate.
type Periods []Period

func (p Periods) Valid() bool {
	if len(p) == 0 {
		return false
	}
	for _, elem := range p {
		if !elem.Valid() {
			return false
		}
	}
	return true
}

func (p Periods) Start() time.Time {
	if len(p) == 0 {
		return time.Time{}
	}
	min := p[0].Start()
	for i := 1; i < len(p); i++ {
		min = cabiriaTime.Min(min, p[i].Start())
	}
	return min
}

func (p Periods) End() time.Time {
	if len(p) == 0 {
		return time.Time{}
	}
	max := p[0].End()
	for i := 1; i < len(p); i++ {
		max = cabiriaTime.Max(max, p[i].End())
	}
	return max
}

func (p Periods) TransformToNew(start, end time.Time) Period {
	// Determine bounds of many
	manyMin := p.Start()
	manyMax := p.End()

	// Shift to min
	dist := start.Sub(manyMin)
	shifted := shiftPeriods(p, dist)

	// Scale to max
	manySpan := manyMax.Sub(manyMin)
	desiredSpan := end.Sub(start)
	scaleFactor := float64(desiredSpan) / float64(manySpan)
	scaled := scalePeriods(shifted, start, scaleFactor)

	return scaled
}

func scalePeriods(p Periods, origin time.Time, factor float64) Periods {
	var results []Period
	for _, elem := range p {
		results = append(results, Scale(elem, origin, factor))
	}
	return results
}

func shiftPeriods(periods Periods, amount time.Duration) Periods {
	var results []Period
	for _, elem := range periods {
		results = append(results, Shift(elem, amount))
	}
	return results
}

func FixOverlaps(many Periods) Periods {
	if len(many) == 0 {
		return []Period{}
	}
	Sort(many)
	var result []Period
	currentSet := Periods([]Period{many[0]})
	for i := 1; i < len(many); i++ {
		elem := many[i]
		if DoesOverlap(elem, currentSet) {
			currentSet = append(currentSet, elem)
		} else {
			result = append(result, seperate(currentSet)...)
			currentSet = Periods([]Period{elem})
		}
	}
	result = append(result, seperate(currentSet)...)
	return result
}

func MergeTouching(many Periods, mergeFunc func(a, b Period) Period) Periods {
	if len(many) == 0 {
		return []Period{}
	}
	wrappedMergeFunc := func(set Periods) Periods {
		return Periods([]Period{merge(set, mergeFunc)})
	}
	Sort(many)
	var result []Period
	currentSet := Periods([]Period{many[0]})
	for i := 1; i < len(many); i++ {
		elem := many[i]
		if Touching(elem, currentSet) {
			currentSet = append(currentSet, elem)
		} else {
			result = append(result, wrappedMergeFunc(currentSet)...)
			currentSet = Periods([]Period{elem})
		}
	}
	result = append(result, wrappedMergeFunc(currentSet)...)
	return result
}

func CoverGaps(many Periods) Periods {
	result := make(Periods, len(many))
	copy(result, many)
	Sort(result)
	for i := 0; i < len(result)-1; i++ {
		// For each pair
		before := result[i]
		after := result[i+1]
		// If there is a gap
		if !Touching(before, after) {
			dist := after.Start().Sub(before.End())
			// Figure out what proportion of the gap goes to each
			beforeDuration := Duration(before)
			afterDuration := Duration(after)
			totalDuration := beforeDuration + afterDuration
			gapToAfter := time.Duration(float64(dist) * (float64(afterDuration) / float64(totalDuration)))
			meetingPoint := after.Start().Add(-gapToAfter)

			// Modify periods
			before = before.TransformToNew(before.Start(), meetingPoint)
			after = after.TransformToNew(meetingPoint, after.End())

			// Store result
			result[i] = before
			result[i+1] = after
		}
	}
	return result
}

func Sort(many Periods) {
	sort.Slice(many, func(i, j int) bool {
		return many[i].Start().Before(many[j].Start())
	})
}

func seperate(many Periods) Periods {
	var results Periods
	spanDuration := float64(Duration(many))
	overlappingDuration := float64(durationSum(many))
	origin := many.Start()
	newStart := origin
	for _, elem := range many {
		percentageOfTotal := float64(Duration(elem)) / overlappingDuration
		newEnd := newStart.Add(time.Duration(spanDuration * percentageOfTotal))
		results = append(results, elem.TransformToNew(newStart, newEnd))
		newStart = newEnd
	}
	return results
}

func merge(many Periods, mergeFunc func(a, b Period) Period) Period {
	if len(many) == 0 {
		return nil
	}
	base := many[0]
	for i := 1; i < len(many); i++ {
		base = mergeFunc(base, many[i])
	}
	return base
}

func durationSum(many Periods) time.Duration {
	sum := time.Duration(0)
	for _, elem := range many {
		sum += Duration(elem)
	}
	return sum
}
