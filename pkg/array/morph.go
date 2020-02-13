package array

// CloseBoolArray will close any "gaps" in the array if the gap is smaller
//  than the threshold.
//  threshold must be larger than 1, otherwise the items will be unchanged.
func CloseBoolArray(items []bool, threshold uint) {
	lastOn := len(items)
	for currentOn, isOn := range items {
		if isOn {
			if needsToBeClosed(lastOn, currentOn, threshold) {
				close(items, lastOn, currentOn)
			}
			lastOn = currentOn
		}
	}
}

func needsToBeClosed(lastOn, currentOn int, threshold uint) bool {
	dist := currentOn - lastOn
	return dist > 1 && dist <= int(threshold)
}

func close(items []bool, lastOn, currentOn int) {
	for i := lastOn + 1; i < currentOn; i++ {
		items[i] = true
	}
}
