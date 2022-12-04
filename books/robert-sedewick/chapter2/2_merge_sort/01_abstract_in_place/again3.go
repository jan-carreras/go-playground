package abstract_in_place

import "golang.org/x/exp/constraints"

// Fucking repeat it again, because I'm not fluent enough

type Sort3[T constraints.Ordered] struct {
	aux []T
}

func (s *Sort3[T]) Sort(input []T) {
	s.aux = make([]T, len(input))

	s.sort(input, 0, len(input)/2, len(input))

	s.aux = nil
}

func (s *Sort3[T]) sort(input []T, lo int, mid int, hi int) {
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
