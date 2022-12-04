package bottom_up

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"io"
)

type BottomUpSortRepeat[T constraints.Ordered] struct {
	aux    []T
	output io.Writer
}

func (s *BottomUpSortRepeat[T]) WithDebug(output io.Writer) *BottomUpSortRepeat[T] {
	s.output = output
	return s
}

func (s *BottomUpSortRepeat[T]) Sort(input []T) {
	s.aux = make([]T, len(input))
	N := len(input)

	for sz := 1; sz < N; sz *= 2 {
		for lo := 0; lo < N-sz; lo += sz + sz {

		}
	}
}

func (s *BottomUpSortRepeat[T]) merge(input []T, lo int, mid int, hi int) {
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

func (s *BottomUpSortRepeat[T]) debug(str string) {
	if s.output != nil {
		fmt.Fprint(s.output, str)
	}
}
