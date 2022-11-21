package main

import (
	"sort"
)

// 1.1.29 Equal keys. Add to BinarySearch a static method rank() that takes a key
// and a sorted array of int values (some of which may be equal) as arguments and
// returns the number of elements that are smaller than the key and a similar
// method count() that returns the number of elements equal to the key. Note : If
// i and j are the values returned by rank(key, a) and count(key, a)
// respectively, then a[i..i+j-1] are the values in the array that are equal to
// key.

type BinarySearch struct {
	lst []int
}

func NewBinarySearch(lst []int) BinarySearch {
	sort.Ints(lst)
	return BinarySearch{lst: lst}

}

// Search returns the index of the element being searched
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

// SmallerThan returns the number of elements that are smaller than the key
func (b BinarySearch) SmallerThan(n int) int {
	i := b.Search(n)
	if i == -1 {
		return i
	}

	min, _ := b.minHi(i)
	return min
}

// Count returns the number of elements equal to the key
func (b BinarySearch) Count(n int) int {
	i := b.Search(n)
	if i == -1 {
		return i
	}

	min, hi := b.minHi(i)
	return hi - min + 1
}

// minHi returns the "lo" and "hi" indexes of all the elements of equal value on lst[idx]
func (b BinarySearch) minHi(idx int) (lo int, hi int) {
	lo, hi = idx, idx
	for lo > 1 && b.lst[lo-1] == b.lst[idx] {
		lo--
	}

	for hi < len(b.lst)-1 && b.lst[hi+1] == b.lst[hi] {
		hi++
	}

	return lo, hi
}
