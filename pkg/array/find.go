package array

// Contains returns true if search is in items, otherwise false.
func Contains(items []string, search string) bool {
	for _, item := range items {
		if item == search {
			return true
		}
	}
	return false
}
