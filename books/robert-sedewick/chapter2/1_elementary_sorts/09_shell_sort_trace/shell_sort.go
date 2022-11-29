package shell_trace

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"io"
)

// 2.1.9 Show,in the style of the example trace with Algorithm2.3, how shell sort
// the array:
//
// E A S Y S H E L L S O R T Q U E S T I O N.

type ShellSort[T constraints.Ordered] struct {
	writer io.Writer
}

func (s *ShellSort[T]) WithWriter(w io.Writer) *ShellSort[T] {
	s.writer = w
	return s
}

// Sort sorts the input
func (s *ShellSort[T]) Sort(input []T) {
	N := len(input)
	h := 1
	for h < N/3 {
		h = 3*h + 1 // 1, 4, 13, 40, 121, 364, 1093, ...
	}

	//s.write(fmt.Sprintf("N=%d h=%d\n", N, h))

	for h >= 1 {
		//s.write(fmt.Sprintf("h-start=%d\n", h))
		for i := h; i < N; i++ {
			for j := i; j >= h && input[j] < input[j-1]; j -= h {
				s.write(fmt.Sprintf("i=%2d j=%2d h=%2d %v\n", i, j, h, input))
				input[j], input[j-1] = input[j-1], input[j]
			}
		}
		h /= 3 // ????
		//s.write(fmt.Sprintf("h-end=%d\n", h))
	}
}

func (s *ShellSort[T]) write(str string) {
	if s.writer != nil {
		fmt.Fprint(s.writer, str)
	}
}
