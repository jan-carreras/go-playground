package sorted_queues

import (
	"container/list"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMergeQueues(t *testing.T) {
	t.Run("q2 larger than q1", func(t *testing.T) {
		q1 := makeQueue([]int{1, 3, 7})
		q2 := makeQueue([]int{2, 5, 9, 10, 16, 20})
		result := mergeQueues(q1, q2)
		require.Equal(t, []int{1, 2, 3, 5, 7, 9, 10, 16, 20}, getElements(result))
	})

	t.Run("q1 larger than q1", func(t *testing.T) {
		q1 := makeQueue([]int{2, 5, 9, 10, 16, 20})
		q2 := makeQueue([]int{1, 3, 7})
		result := mergeQueues(q1, q2)
		require.Equal(t, []int{1, 2, 3, 5, 7, 9, 10, 16, 20}, getElements(result))
	})

	t.Run("same size queues", func(t *testing.T) {
		q1 := makeQueue([]int{2, 5, 9})
		q2 := makeQueue([]int{1, 3, 7})
		result := mergeQueues(q1, q2)
		require.Equal(t, []int{1, 2, 3, 5, 7, 9}, getElements(result))
	})

}

func TestMergeQueues_OneQueueIsEmpty(t *testing.T) {

	t.Run("q1 is empty", func(t *testing.T) {
		q1 := makeQueue([]int{1, 2, 3})
		q2 := makeQueue([]int{})
		result := mergeQueues(q1, q2)
		require.Equal(t, []int{1, 2, 3}, getElements(result))
	})

	t.Run("q2 is empty", func(t *testing.T) {
		q1 := makeQueue([]int{})
		q2 := makeQueue([]int{1, 2, 3})
		result := mergeQueues(q1, q2)
		require.Equal(t, []int{1, 2, 3}, getElements(result))
	})

}

func makeQueue(input []int) *list.List {
	l := list.New()
	for _, i := range input {
		l.PushBack(i)
	}
	return l
}

func getElements(q *list.List) []int {
	ints := make([]int, 0, q.Len())
	n := q.Front()
	for n != nil {
		ints = append(ints, n.Value.(int))
		n = n.Next()
	}
	return ints
}
