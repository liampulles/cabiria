package math

import (
	"fmt"
	"math"
)

func EuclideanDistance(a []float64, b []float64) (float64, error) {
	if a == nil {
		return -1.0, fmt.Errorf("a may not be nil")
	}
	if b == nil {
		return -1.0, fmt.Errorf("b may not be nil")
	}
	if len(a) != len(b) {
		return -1.0,
			fmt.Errorf("a and b have differet lengths. a length: %d, b length: %d",
				len(a), len(b))
	}
	total := 0.0
	for i, ai := range a {
		bi := b[i]
		diff := ai - bi
		total += diff * diff
	}
	return math.Sqrt(total), nil
}
