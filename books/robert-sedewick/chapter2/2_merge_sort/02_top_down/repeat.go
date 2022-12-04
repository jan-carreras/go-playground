package top_down

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"io"
)

type TopDownRepeat[T constraints.Ordered] struct {
	aux    []T
	output io.Writer
}

func (t *TopDownRepeat[T]) Sort(input []T) {
	t.aux = make([]T, len(input))
	t.sort(input, 0, len(input)-1)

}

func (t *TopDownRepeat[T]) sort(input []T, lo int, hi int) {
	if hi >= lo {
		return
	}

	mid := lo + (hi-lo)/2
	t.sort(input, lo, mid)
	t.sort(input, mid+1, hi)
	t.merge(input, lo, mid, hi)
}

func (t *TopDownRepeat[T]) merge(input []T, lo int, mid int, hi int) {
	for k := lo; k <= hi; k++ {
		t.aux[k] = input[k]
	}

	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			copy(input[k:], t.aux[j:])
			break
		} else if j > hi {
			copy(input[k:], t.aux[i:mid+1])
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

func (t *TopDownRepeat[T]) debug(str string) {
	if t.output != nil {
		fmt.Fprint(t.output, str)
	}
}

func (t *TopDownRepeat[T]) WithDebug(output io.Writer) *TopDownRepeat[T] {
	t.output = output
	return t

}
