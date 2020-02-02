package array

func Contains(items []string, search string) bool {
	for _, item := range items {
		if item == search {
			return true
		}
	}
	return false
}
