package period

import (
	"sort"
	"time"

	cabiriaTime "github.com/liampulles/cabiria/pkg/time"
)

// Periods is a slice of Period. It can itself be considered a Period (and we
//  implement Period for Periods)... see below.
type Periods []Period

// Valid is true for Periods when there is at least one element, and all elements
//  are themselves Valid.
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

// Start is the minimum start of all elements in Periods.
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

// End is the maximum end of all elements in Periods.
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

// TransformToNew scales and shifts the elements of Periods, such that
//  the minimum start of all elements is now "start", and the maximum end
//  of all elements is now "end". The elements are also transformed into new
//  variants, and the relative relationship between elements is unchanged -
//  e.g. if elements 2 and 7 were overlapping, they will continue to overlap by
//  the same percentage after the transformation.
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

// FixOverlaps will adjust any set of overlapping elements in many such that
//  their bounds touch, and they share the span of their overlapping set in
//  proportion to their original Durations.
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
			result = append(result, separate(currentSet)...)
			currentSet = Periods([]Period{elem})
		}
	}
	result = append(result, separate(currentSet)...)
	return result
}

// MergeTouching will merge any touching periods using mergeFunc.
//  mergeFunc should return a period which has Start = a.Start() and end
//  = b.End(), otherwise the result is not guaranteed to have non-touching
//  elements.
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

// CoverGaps will close any gaps between close elements by stretching those
//  elements to cover the gap. The degree to which the elements are stretched is
//  determined by their original Duration.
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

// Sort orders the elements naturally.
func Sort(many Periods) {
	sort.Slice(many, func(i, j int) bool {
		if many[i].Start().Equal(many[j].Start()) {
			return many[i].End().Before(many[j].End())
		}
		return many[i].Start().Before(many[j].Start())
	})
}

func separate(many Periods) Periods {
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
