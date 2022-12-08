package natural_mergesort

import (
	"golang.org/x/exp/constraints"
)

// 2.2.16 Natural mergesort.
//
// Write a version of bottom-up mergesort that takes advantage of order in the
// array by proceeding as follows each time it needs to find two arrays to merge:
// find a sorted subarray (by incrementing a pointer until finding an entry that
// is smaller than its predecessor in the array), then find the next, then merge
// them. Analyze the running time of this algorithm in terms of the array size
// and the number of maximal increasing sequences in the array.

func Sort[T constraints.Ordered](input []T) {
	aux := make([]T, len(input))
	if len(input) == 0 {
		return
	}

	N := len(input)

LOOP:
	for {
		for lo := 0; lo < N; lo++ {
			mid := sortedUntil(input, lo)
			if mid == N-1 && lo == 0 {
				break LOOP // Sorted array, we're done
			}
			hi := sortedUntil(input, mid+1)
			merge(input, aux, lo, mid, hi)
			lo = hi
		}
	}
}

func sortedUntil[T constraints.Ordered](input []T, idx int) (mid int) {
	for mid = idx; mid < len(input)-1; mid++ {
		if input[mid] > input[mid+1] {
			return mid
		}
	}
	return mid
}

func merge[T constraints.Ordered](input, aux []T, lo, mid, hi int) {
	if hi >= len(input) {
		return
	}

	for k := lo; k <= hi; k++ {
		aux[k] = input[k]
	}

	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			input[k] = aux[j]
			j++
		} else if j > hi {
			input[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			input[k] = aux[j]
			j++
		} else {
			input[k] = aux[i]
			i++
		}
	}
}
