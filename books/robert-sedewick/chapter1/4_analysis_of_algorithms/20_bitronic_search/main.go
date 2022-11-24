package bitronic_search

import (
	"sort"
)

// 1.4.20 Bitonic search.
//
// An array is bitonic if it comprises an increasing sequence of integers
// followed immediately by a decreasing sequence of integers. Write a program
// that, given a bitonic array of N distinct int values, determines whether a
// given integer is in the array. Your program should use ~3lg N compares in the
// worst case.

// Ideas: find the "tip", and from [0,tip] -> classic binary search
// 							  from [tip+1, len()] -> opposite binary search
// N + log(N)

// Idea2: Fucking sort the damn input, regardless of its ordering? But not really, because it takes N logN already

// Idea3: Yeah, tip + normal binary search + reverse binary search should be ~3lg N

// What if I ignore the "search for the tip"

func BitronicSearch(input []int, search int) bool {
	if len(input) < 3 {
		// Is impossible to have an ascending-descending list of integer without
		// at least 3 numbers
		panic("not a bitronic input")
	}

	tip := findTip(input)

	if binarySearch(input, search, 0, tip) {
		return true
	}

	return reverseSearch(input, search, tip+1, len(input)-1)
}

func binarySearch(input []int, search, lo, hi int) bool {
	for lo <= hi {
		mid := lo + (hi-lo)/2
		switch {
		case search > input[mid]:
			lo = mid + 1
		case search < input[mid]:
			hi = mid - 1
		default:
			return true
		}
	}

	return false
}

func reverseSearch(input []int, search, lo, hi int) bool {
	for lo <= hi {
		mid := lo + (hi-lo)/2
		switch {
		case search < input[mid]:
			lo = mid + 1
		case search > input[mid]:
			hi = mid - 1

		default:
			return true
		}
	}

	return false
}

func findTip(input []int) (tip int) {
	lo, hi := 1, len(input)-2
	for lo <= hi {
		mid := lo + (hi-lo)/2
		prev, node, next := input[mid-1], input[mid], input[mid+1]
		switch {
		case prev < node && node < next: // Ascending
			lo = mid + 1
		case prev > node && node > next: // Descending
			hi = mid - 1
		case prev < node && node > next: // The tip
			tip = mid
			return mid
		default:
			panic("yo, wtf")
		}
	}
	return tip
}

// Complexity NlogN + LogN
func BitronicSearchSorting(input []int, search int) bool {
	sort.Ints(input)

	lo, hi := 0, len(input)-1
	for lo > hi {
		mid := lo + ((hi - lo) / 2)
		switch {
		case search > input[mid]:
			lo = mid + 1
		case search < input[mid]:
			hi = mid - 1
		default:
			return true
		}
	}

	return false
}

// Complexity: N
func BitronicSearchBruteforce(input []int, search int) bool {
	for i := 0; i < len(input); i++ {
		if input[i] == search {
			return true
		}
	}

	return false
}
