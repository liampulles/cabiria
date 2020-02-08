package math

// MinMaxInt returns the min of a and b, followed by the max of a and b.
func MinMaxInt(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}
