package __1_28_remove_duplicated

import "sort"

// Remove duplicates. Modify the test client in BinarySearch to remove any
// duplicate keys in the whitelist after the sort.

type BinarySearch struct {
	lst []int
}

func NewBinarySearch(lst []int) BinarySearch {
	if len(lst) == 0 {
		return BinarySearch{lst: lst}
	}

	// Sorting
	sort.Ints(lst)

	// Removing duplicates
	cleanLst := []int{lst[0]}
	for i := 1; i < len(lst)-1; i++ {
		if lst[i] == lst[i-1] {
			continue
		}
		cleanLst = append(cleanLst, lst[i])
	}

	return BinarySearch{lst: cleanLst}
}

func (b BinarySearch) Search(n int) int {
	lo, hi := 0, len(b.lst)-1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		switch {
		case n > b.lst[mid]:
			lo = mid + 1
		case n < b.lst[mid]:
			hi = mid - 1
		default:
			return mid
		}
	}

	return -1
}

func NewBinarySearch_INCORRECT(lst []int) BinarySearch {
	// This is INCORRECT because it's BEFORE the sort!!
	set := make(map[int]bool)
	for _, i := range lst {
		set[i] = true
	}

	lst = make([]int, 0, len(set))
	for k := range set {
		lst = append(lst, k)
	}

	sort.Ints(lst)
	bs := BinarySearch{lst: lst}

	return bs
}
