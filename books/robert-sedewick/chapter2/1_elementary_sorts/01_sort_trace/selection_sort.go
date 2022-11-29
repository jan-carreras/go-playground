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
	writer io.Writer
}

func (s *SelectionSort[T]) WithWriter(w io.Writer) *SelectionSort[T] {
	s.writer = w
	return s
}

// Sort sorts the input
func (s *SelectionSort[T]) Sort(input []T) {
	N := len(input)
	for i := 0; i < N; i++ {
		min := i
		for j := i + 1; j < N; j++ {
			if input[i] > input[j] {
				min = j
			}
		}

		s.write(fmt.Sprintf("i=%2d min=%2d %v\n", i, min, input))
		input[i], input[min] = input[min], input[i]
	}
}

func (s *SelectionSort[T]) write(str string) {
	if s.writer != nil {
		fmt.Fprint(s.writer, str)
	}
}
