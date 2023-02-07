package binarysearch

func SearchInts(tree []int, search int) int {
	lower, upper := 0, len(tree)-1

	for lower <= upper {
		middle := lower + ((upper - lower) / 2)
		if tree[middle] == search {
			return middle
		} else if tree[middle] > search {
			upper = middle - 1
		} else {
			lower = middle + 1
		}
	}

	return -1
}
