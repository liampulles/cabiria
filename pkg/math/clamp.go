package math

func Clamp(val float64, min float64, max float64) float64 {
	if val <= min {
		return min
	}
	if val >= max {
		return max
	}
	return val
}
