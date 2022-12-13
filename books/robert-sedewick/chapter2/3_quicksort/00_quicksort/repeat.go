package quicksort

import (
	"golang.org/x/exp/constraints"
	"math/rand"
)

func Sort2[T constraints.Ordered](input []T) {
	// Fisher and Yates randomization
	for i := len(input) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		input[i], input[j] = input[j], input[i]
	}

	sort2(input, 0, len(input)-1)
}

func sort2[T constraints.Ordered](input []T, lo int, hi int) {
	if lo >= hi {
		return
	}

	p := partition2(input, lo, hi)
	sort2(input, lo, p-1)
	sort2(input, p+1, hi)
}

func partition2[T constraints.Ordered](input []T, lo int, hi int) int {
	pivot := input[hi]

	i := lo
	for j := lo; j < hi; j++ {
		if input[j] < pivot {
			input[i], input[j] = input[j], input[i]
			i++
		}
	}

	input[i], input[hi] = input[hi], input[i]

	return i
}
