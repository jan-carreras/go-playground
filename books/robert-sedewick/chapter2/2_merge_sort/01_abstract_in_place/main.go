package abstract_in_place

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"io"
)

type Sort[T constraints.Ordered] struct {
	aux    []T
	output io.Writer
}

// InPlace returns a
func (s *Sort[T]) InPlace(input []T) {
	s.aux = make([]T, len(input))

	lo, hi := 0, len(input)
	mid := lo + (hi-lo)/2
	// Merge from [0, len(input)] -> both included
	s.inPlaceMerge(input, lo, mid, hi)

	s.aux = nil // Liberate resources
}

// inPlaceMerge merge two slices into one
func (s *Sort[T]) inPlaceMerge(input []T, lo, mid, hi int) {
	i, j := lo, mid

	for k := lo; k < hi; k++ {
		s.aux[k] = input[k]
	}

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
		s.debug(fmt.Sprintf("%v\n", input[:k+1]))
	}
}

func (s *Sort[T]) debug(str string) {
	if s.output != nil {
		fmt.Fprint(s.output, str)
	}
}

func (s *Sort[T]) WithDebug(output io.Writer) *Sort[T] {
	s.output = output
	return s
}
