package linked_lists

import (
	"container/list"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMergeSort(t *testing.T) {
	input := []int{100, 200, 300, 400, 120, 130, 140, 250, 500, 520, 10, 20, 30, 900}
	l := newList(input)

	MergeSort(l)
	require.Equal(t, []int{10, 20, 30, 100, 120, 130, 140, 200, 250, 300, 400, 500, 520}, toInt(l))
}

func TestMerge(t *testing.T) {
	input := []int{100, 200, 300, 400, 120, 130, 140, 250, 500, 520, 10, 20, 30, 900}
	l := newList(input)

	lo := l.Front()
	mid := findIncreasingBlock(lo)
	hi := findIncreasingBlock(mid.Next())

	merge(l, lo, mid, hi)

	require.Equal(t, []int{100, 120, 130, 140, 200, 250, 300, 400, 500, 520, 10, 20, 30}, toInt(l))

}

func newList(input []int) *list.List {
	l := list.New()
	for _, value := range input {
		l.PushBack(value)
	}
	return l
}

func TestFindIncreasingBlock(t *testing.T) {
	input := []int{100, 200, 300, 400, 120, 130, 140, 250, 500, 520, 10, 20, 30, 900}
	l := newList(input)

	n := findIncreasingBlock(l.Front())
	require.NotNil(t, n)
	require.EqualValues(t, 400, n.Value)

	n = findIncreasingBlock(n.Next())
	require.NotNil(t, n)
	require.EqualValues(t, 520, n.Value)

	n = findIncreasingBlock(n.Next())
	require.NotNil(t, n)
	require.EqualValues(t, 900, n.Value)

	n = findIncreasingBlock(n.Next())
	require.Nil(t, n)
}
