package array

import (
	"github.com/liampulles/cabiria/pkg/math"
)

// Opening is the morphological opening
func Opening(set []bool, kernel uint) []bool {
	set = Dilation(set, kernel)
	set = Erosion(set, kernel)
	return set
}

// Closing is the morphological closing
func Closing(set []bool, kernel uint) []bool {
	set = Erosion(set, kernel)
	set = Dilation(set, kernel)
	return set
}

// Erosion is the morphological erosion
func Erosion(set []bool, kernel uint) []bool {
	if kernel == 0 {
		return set
	}
	clone := clone(set)
	start := int(kernel / 2)
	end := int(kernel) - start
	for i := 0; i < len(set); i++ {
		if allForeground(set, i-start, i+end) {
			// Do nothing
		}
		if anyBackground(set, i-start, i+end) {
			clone[i] = false
		}
	}
	return clone
}

// Dilation is the morphological dilation
func Dilation(set []bool, kernel uint) []bool {
	if kernel == 0 {
		return set
	}
	clone := clone(set)
	start := int(kernel / 2)
	end := int(kernel) - start
	for i := 0; i < len(set); i++ {
		if anyForeground(set, i-start, i+end) {
			clone[i] = true
		}
		if allCorrespondingBackground(set, start-i, i+end) {
			clone[i] = false
		}
	}
	return clone
}

func allForeground(set []bool, start int, end int) bool {
	if start < 0 || end > len(set) {
		return false
	}
	for i := start; i < end; i++ {
		if !set[i] {
			return false
		}
	}
	if (end - start) == 0 {
		return false
	}
	return true
}

func allCorrespondingBackground(set []bool, start int, end int) bool { //?
	start = math.ClampInt(start, 0, len(set))
	end = math.ClampInt(end, 0, len(set))
	for i := start; i < end; i++ {
		if set[i] {
			return false
		}
	}
	if (end - start) == 0 {
		return false
	}
	return true
}

func anyForeground(set []bool, start int, end int) bool {
	start = math.ClampInt(start, 0, len(set))
	end = math.ClampInt(end, 0, len(set))
	for i := start; i < end; i++ {
		if set[i] {
			return true
		}
	}
	return false
}

func anyBackground(set []bool, start int, end int) bool {
	if start < 0 || end > len(set) {
		return true
	}
	for i := start; i < end; i++ {
		if !set[i] {
			return true
		}
	}
	return false
}

func clone(set []bool) []bool {
	clone := make([]bool, len(set))
	copy(clone, set)
	return clone
}
