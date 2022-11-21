package __0_0_binary_search

func Search(lst []int, n int) int {
	lo, mid, hi := 0, len(lst)/2, len(lst)-1

	for lo <= hi {
		mid = lo + ((hi - lo) / 2)
		switch {
		case n > lst[mid]:
			lo = mid + 1
		case n < lst[mid]:
			hi = mid - 1
		default:
			return mid
		}

	}

	return -1
}
