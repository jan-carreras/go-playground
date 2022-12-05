package improvements

import (
	"golang.org/x/exp/constraints"
)

// 2.2.11 Improvements. Implement the three improvements to mergesort that are
// described in the text on page 275:
//
// - add a cutoff for small subarrays,
// - test whether the array is already in order,
// - and avoid the copy by switching arguments in the recursive code.

var cutoff = 6

func Sort[T constraints.Ordered](input []T) {
	if isSorted(input) {
		return
	}

	aux := make([]T, len(input))
	copy(aux, input)
	sort(input, aux, 0, len(input)-1, false)
}

func sort[T constraints.Ordered](input, aux []T, lo, hi int, swapInputs bool) {
	if hi <= lo {
		return
	}

	if swapInputs {
		input, aux = aux, input
	}

	// add a cutoff for small subarrays
	if hi-lo <= cutoff {
		insertSort(input, lo, hi)
		return
	}

	mid := lo + (hi-lo)/2
	sort(input, aux, lo, mid, !swapInputs)
	sort(input, aux, mid+1, hi, !swapInputs)
	merge(input, aux, lo, mid, hi)
}

func insertSort[T constraints.Ordered](input []T, lo, hi int) {
	// Insert sort
	for i := lo; i <= hi; i++ {
		for j := i; j > 0 && input[j] < input[j-1]; j-- {
			input[j], input[j-1] = input[j-1], input[j]
		}
	}
}

func merge[T constraints.Ordered](input []T, aux []T, lo, mid, hi int) {
	// and avoid the copy by switching arguments in the recursive code
	/*	for k := lo; k <= hi; k++ {
		aux[k] = input[k]
	}*/

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

func isSorted[T constraints.Ordered](input []T) bool {
	// test whether the array is already in order
	for i := 1; i < len(input); i++ {
		if input[i] < input[i-1] {
			return false
		}
	}
	return true
}
