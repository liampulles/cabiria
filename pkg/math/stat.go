package math

import (
	"fmt"
)

// Add performs a vector additon of a and b. If a and b differ in size, an
// error is returned.
func Add(a []float64, b []float64) ([]float64, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("a and b need to have same length. A length: %d, b length: %d", len(a), len(b))
	}
	sum := make([]float64, len(a))
	for i, ai := range a {
		sum[i] = ai + b[i]
	}
	return sum, nil
}

// Square performs an element by element square of vector a,
// i.e. the dot product of a with itself.
func Square(a []float64) []float64 {
	sum := make([]float64, len(a))
	for i, ai := range a {
		sum[i] = ai * ai
	}
	return sum
}
