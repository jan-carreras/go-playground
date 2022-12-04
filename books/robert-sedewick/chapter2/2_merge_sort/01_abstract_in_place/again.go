package abstract_in_place

import (
	"golang.org/x/exp/constraints"
)

type Sort2[T constraints.Ordered] struct {
	aux []T
}

func (s *Sort2[T]) Sort(input []T) {
	// create aux
	s.aux = make([]T, len(input))

	s.sort(input, 0, len(input)/2, len(input))

	s.aux = nil // Clean memory
}

func (s *Sort2[T]) sort(input []T, lo, mid, hi int) {
	for k := lo; k < hi; k++ {
		s.aux[k] = input[k]
	}

	i, j := lo, mid
	for k := lo; k < hi; k++ {
		if i >= mid {
			input[k] = s.aux[j]
			j++
		} else if j >= hi {
			input[k] = s.aux[i]
			i++
		} else if s.aux[j] < s.aux[i] {
			input[k] = s.aux[j]
			j++
		} else {
			input[k] = s.aux[i]
			i++
		}
	}
}
