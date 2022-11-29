package insertion_trace

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"io"
)

// 2.1.4 Show, in the style of the example trace with Algorithm2.2, how insertion
// sort the array:
//
// E A S Y Q U E S T I O N

type InsertSort[T constraints.Ordered] struct {
	writer io.Writer
}

func (s *InsertSort[T]) WithWriter(w io.Writer) *InsertSort[T] {
	s.writer = w
	return s
}

// Sort sorts the input
func (s *InsertSort[T]) Sort(input []T) {
	N := len(input)
	for i := 1; i < N; i++ {
		for j := i; j > 0 && input[j] < input[j-1]; j-- {
			s.write(fmt.Sprintf("i=%2d j=%2d %v\n", i, j, input))
			input[j], input[j-1] = input[j-1], input[j]
		}
	}
}

func (s *InsertSort[T]) write(str string) {
	if s.writer != nil {
		fmt.Fprint(s.writer, str)
	}
}
