package sublinear_extra_space

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// NOTE: Not finished and probably wrong. I'm using M*2 space, rather than M :/

// 2.2.12 Sublinear extra space. Develop a merge implementation that reduces the
// extra space requirement to max(M, N/M), based on the following idea: Divide
// the array into N/M blocks of size M (for simplicity in this description,
// assume that N is a multiple of M). Then, (i) considering the blocks as items
// with their first key as the sort key, sort them using selection sort; and (ii)
// run through the array merging the first block with the second, then the second
// block with the third, and so forth.

// M : Arbitrary value - could be parametrised on the MergeSort function
const M = 3

func MergeSort[T constraints.Ordered](input []T) {
	if len(input)%M != 0 {
		panic(fmt.Errorf("pre-condition not met. input must be divisible by %d", M))
	}

	// Not mentioned in the exercise, but each block must be sorted itself
	insertSortEachBlock(input)

	// (i) considering the blocks as items with their first key as the sort key, sort
	// them using selection sort
	sortByBlock(input)

	// (ii) run through the array merging the first block with the second, then the
	// second block with the third, and so forth
	mergeBlocks(input)
}

func sortByBlock[T constraints.Ordered](input []T) {
	// Iterate thru all the Sorting Keys of each group
	for i := 0; i < len(input); i += M {
		min := i
		for j := i; j < len(input); j += M {
			if input[j] < input[min] {
				min = j
			}
		}

		// Swap two groups
		for s := 0; s < M; s++ {
			input[i+s], input[s+min] = input[s+min], input[i+s]
		}
	}
}

func mergeBlocks[T constraints.Ordered](input []T) {
	// Develop a merge implementation that reduces the extra space requirement to max(M, N/M)
	// NOTE: This program is running with M*2, so it's not correct :/
	aux := make([]T, M*2)
	for i := 0; i+M < len(input); i += M {
		lo, mid, hi := i, i+M-1, i+M+M-1
		merge(input, aux, lo, mid, hi)
	}
}

func merge[T constraints.Ordered](input, aux []T, lo, mid, hi int) {
	for k := lo; k <= hi; k++ {
		aux[k-lo] = input[k]
	}

	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			input[k] = aux[j-lo]
			j++
		} else if j > hi {
			input[k] = aux[i-lo]
			i++
		} else if aux[j-lo] < aux[i-lo] {
			input[k] = aux[j-lo]
			j++
		} else {
			input[k] = aux[i-lo]
			i++
		}
	}
}

// Not needed for anything, afaik — just playing around
func insertSortEachBlock[T constraints.Ordered](input []T) {
	for k := 0; k < len(input); k += M {
		lo, hi := k, k+M-1

		for i := lo; i <= hi; i++ {
			min := i
			for j := i; j <= hi; j++ {
				if input[j] < input[i] {
					min = j
				}
			}

			// Swapping values
			input[i], input[min] = input[min], input[i]
		}
	}
}
