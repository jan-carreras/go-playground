package aux

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"io"
)

type TopDown[T constraints.Ordered] struct {
	output io.Writer
}

func (t *TopDown[T]) Sort(input []T) {
	aux := make([]T, len(input))
	t.sort(input, aux, 0, len(input)-1)

}

func (t *TopDown[T]) sort(input []T, aux []T, lo int, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	t.sort(input, aux, lo, mid)
	t.sort(input, aux, mid+1, hi)
	t.merge(input, aux, lo, mid, hi)
}

func (t *TopDown[T]) merge(input []T, aux []T, lo int, mid int, hi int) {
	for k := lo; k <= hi; k++ { // Copy a[lo..hi] to aux[lo..hi]
		aux[k] = input[k]
	}

	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			copy(input[k:], aux[j:])
			break
		} else if j > hi {
			copy(input[k:], aux[i:mid+1])
			break
		} else if aux[j] < aux[i] {
			input[k] = aux[j]
			j++
		} else {
			input[k] = aux[i]
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
