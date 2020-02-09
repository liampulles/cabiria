package array

// OpenBoolArray will open any "trivially small" sections in the array if the
//  gap is smaller than the threshold.
//  threshold must be larger than 1, otherwise the items will be unchanged.
func OpenBoolArray(items []bool, threshold uint) {
	lastOff := len(items)
	for currentOff, isOn := range items {
		if !isOn {
			if needsToBeOpened(lastOff, currentOff, threshold) {
				open(items, lastOff, currentOff)
			}
			lastOff = currentOff
		}
	}
}

func needsToBeOpened(lastOff, currentOff int, threshold uint) bool {
	dist := currentOff - lastOff
	return dist > 1 && dist <= int(threshold)
}

func open(items []bool, lastOff, currentOff int) {
	for i := lastOff + 1; i < currentOff; i++ {
		items[i] = false
	}
}
