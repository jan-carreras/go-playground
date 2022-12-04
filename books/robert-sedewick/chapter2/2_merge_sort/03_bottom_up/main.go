package bottom_up

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"io"
)

type BottomUpSort[T constraints.Ordered] struct {
	aux    []T
	output io.Writer
}

func (s *BottomUpSort[T]) WithDebug(output io.Writer) *BottomUpSort[T] {
	s.output = output
	return s
}

func (s *BottomUpSort[T]) Sort(input []T) {
	s.aux = make([]T, len(input))
	N := len(input)

	for sz := 1; sz < N; sz = sz + sz { // sz: subarray size
		for lo := 0; lo < N-sz; lo += sz + sz { // lo: subarray index
			min := lo + sz + sz - 1
			if N-1 < min {
				min = N - 1
			}
			s.merge(input, lo, lo+sz-1, min)
		}

		min := sz + sz - 1
		if N-1 < min {
			min = N - 1
		}
		s.debug(fmt.Sprintf("%v\n", input[0:min]))
	}
}

func (s *BottomUpSort[T]) merge(input []T, lo int, mid int, hi int) {
	for k := lo; k <= hi; k++ {
		s.aux[k] = input[k]
	}

	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			copy(input[k:hi+1], s.aux[j:hi+1])
			break
		} else if j > hi {
			copy(input[k:hi+1], s.aux[i:mid+1])
			break
		} else if s.aux[j] < s.aux[i] {
			input[k] = s.aux[j]
			j++
		} else {
			input[k] = s.aux[i]
			i++
		}
	}
}

func (s *BottomUpSort[T]) debug(str string) {
	if s.output != nil {
		fmt.Fprint(s.output, str)
	}
}
