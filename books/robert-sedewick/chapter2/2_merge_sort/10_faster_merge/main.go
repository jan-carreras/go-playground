package faster_merge

import (
	"golang.org/x/exp/constraints"
)

// 2.2.10 Faster merge. Implement a version of merge() that copies the second
// half of a[] to aux[] in decreasing order and then does the merge back to a[].
// This change allows you to remove the code to test that each of the halves has
// been exhausted from the inner loop. Note: The resulting sort is not stable.

// marge is the classic merge function — it doesn't care if the two parts are sorted
func merge[T constraints.Ordered](input []T, lo, mid, hi int) {
	aux := make([]T, len(input))
	for k := lo; k <= hi; k++ {
		aux[k] = input[k]
	}

	i, j := lo, mid
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

// This is the merges sorted proposed. The two parts must be ordered beforehand for
// it to work.
func merge2[T constraints.Ordered](input []T, lo, mid, hi int) {
	aux := make([]T, len(input))

	for i := lo; i <= mid; i++ {
		aux[i] = input[i]
	}

	for j := mid; j <= hi; j++ {
		aux[j] = input[hi-j+mid]
	}

	i, j := lo, hi
	for k := lo; k <= hi; k++ {
		if aux[j] < aux[i] {
			input[k] = aux[j]
			j--
		} else {
			input[k] = aux[i]
			i++
		}
	}
}
