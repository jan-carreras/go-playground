package top_down

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"io"
)

type TopDown[T constraints.Ordered] struct {
	aux    []T
	output io.Writer
}

func (t *TopDown[T]) Sort(input []T) {
	t.aux = make([]T, len(input))
	t.sort(input, 0, len(input)-1)

}

func (t *TopDown[T]) sort(input []T, lo int, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	t.sort(input, lo, mid)
	t.sort(input, mid+1, hi)
	t.merge(input, lo, mid, hi)
}

func (t *TopDown[T]) merge(input []T, lo int, mid int, hi int) {
	for k := lo; k <= hi; k++ { // Copy a[lo..hi] to aux[lo..hi]
		t.aux[k] = input[k]
	}

	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			input[k] = t.aux[j]
			j++
		} else if j > hi {
			input[k] = t.aux[i]
			i++
		} else if t.aux[j] < t.aux[i] {
			input[k] = t.aux[j]
			j++
		} else {
			input[k] = t.aux[i]
			i++
		}
	}
	t.debug(fmt.Sprintf("%v\n", input[lo:hi+1]))
}

func (t *TopDown[T]) debug(str string) {
	if t.output != nil {
		fmt.Fprint(t.output, str)
	}
}

func (t *TopDown[T]) WithDebug(output io.Writer) *TopDown[T] {
	t.output = output
	return t

}
