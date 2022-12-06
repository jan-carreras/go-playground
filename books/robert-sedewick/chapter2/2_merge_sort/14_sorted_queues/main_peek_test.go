package sorted_queues

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMergeQueuesPeek(t *testing.T) {
	t.Run("q2 larger than q1", func(t *testing.T) {
		q1 := makeQueue([]int{1, 3, 7})
		q2 := makeQueue([]int{2, 5, 9, 10, 16, 20})
		result := mergeQueuesWithPeek(q1, q2)
		require.Equal(t, []int{1, 2, 3, 5, 7, 9, 10, 16, 20}, getElements(result))
	})

	t.Run("q1 larger than q1", func(t *testing.T) {
		q1 := makeQueue([]int{2, 5, 9, 10, 16, 20})
		q2 := makeQueue([]int{1, 3, 7})
		result := mergeQueuesWithPeek(q1, q2)
		require.Equal(t, []int{1, 2, 3, 5, 7, 9, 10, 16, 20}, getElements(result))
	})

	t.Run("same size queues", func(t *testing.T) {
		q1 := makeQueue([]int{2, 5, 9})
		q2 := makeQueue([]int{1, 3, 7})
		result := mergeQueuesWithPeek(q1, q2)
		require.Equal(t, []int{1, 2, 3, 5, 7, 9}, getElements(result))
	})

}

func TestMergeQueues_OneQueueIsEmptyPeek(t *testing.T) {

	t.Run("q1 is empty", func(t *testing.T) {
		q1 := makeQueue([]int{1, 2, 3})
		q2 := makeQueue([]int{})
		result := mergeQueuesWithPeek(q1, q2)
		require.Equal(t, []int{1, 2, 3}, getElements(result))
	})

	t.Run("q2 is empty", func(t *testing.T) {
		q1 := makeQueue([]int{})
		q2 := makeQueue([]int{1, 2, 3})
		result := mergeQueuesWithPeek(q1, q2)
		require.Equal(t, []int{1, 2, 3}, getElements(result))
	})

}
