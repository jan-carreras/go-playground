package selection_sort

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"io"
)

// 2.1.1 Show, in the style of the example trace with Algorithm 2.1, how
// selection sort sorts the array:
//
// E A S Y Q U E S T I O N.

type SelectionSort[T constraints.Ordered] struct {
	input  []T
	writer io.Writer
}

// Sort sorts the input
func (s *SelectionSort[T]) Sort(input []T) {
	s.input = input

	N := len(s.input)
	for i := 0; i < N; i++ {
		min := i
		for j := i + 1; j < N; j++ {
			if s.Less(i, j) > 0 {
				min = j
			}
		}

		s.write(fmt.Sprintf("i=%2d min=%2d %v\n", i, min, s.input))
		s.Swap(i, min)
	}
}

func (s *SelectionSort[T]) write(str string) {
	if s.writer != nil {
		fmt.Fprint(s.writer, str)
	}
}

// Less returns:
//
//	-1: x<y
//	 0: x==y
//	+1 x>y
func (s *SelectionSort[T]) Less(x, y int) int {
	if s.input[x] < s.input[y] {
		return -1
	} else if s.input[x] > s.input[y] {
		return 1
	}
	return 0
}

// Swap interchanges positions of two elements
func (s *SelectionSort[T]) Swap(x, y int) {
	s.input[x], s.input[y] = s.input[y], s.input[x]
}
