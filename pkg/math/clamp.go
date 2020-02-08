package math

// ClampFloat64 limits val to be between min and max, inclusive.
func ClampFloat64(val float64, min float64, max float64) float64 {
	if val <= min {
		return min
	}
	if val >= max {
		return max
	}
	return val
}
