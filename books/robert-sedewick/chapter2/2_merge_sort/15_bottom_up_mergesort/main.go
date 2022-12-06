package bottom_up_mergesort

import (
	"container/list"
	sorted_queues "github.com/jan-carreras/go-playground/books/robert-sedewick/chapter2/2_merge_sort/14_sorted_queues"
	"golang.org/x/exp/constraints"
)

// 2.2.15 Bottom-up queue mergesort. Develop a bottom-up mergesort implementation
// based on the following approach:
//
// 1. Given N items, create N queues, each containing one of the items.
//
// 2. Create a queue of the N queues.
//
// 3. Then repeatedly apply the merging operation of Exercise 2.2.14 to the first
// two queues and reinsert the merged queue at the end
//
// 4. Repeat until the queue of queues contains only one queue.

func BottomUpMergeSortWithQueues[T constraints.Ordered](input []T) {
	if len(input) <= 1 { // NoOp
		return
	}

	// 2. Create a queue of the N queues.
	lists := list.New()

	// 1. Given N items, create N queues, each containing one of the items.
	for _, value := range input {
		listWithOneItem := list.New()
		listWithOneItem.PushBack(value)
		lists.PushBack(listWithOneItem)
	}

	// 3. Then repeatedly apply the merging operation of Exercise 2.2.14 to the first
	// two queues and reinsert the merged queue at the end
	for lists.Len() > 1 {
		mergedLists := sorted_queues.Merge(
			lists.Remove(lists.Front()).(*list.List), // Remove first element and merge with,
			lists.Remove(lists.Front()).(*list.List), // the second element of the lists
		)
		lists.PushBack(mergedLists)
	}

	finalSortedQueue := lists.Front().Value.(*list.List)

	// Define all values on the input to be the same as the Queue
	node := finalSortedQueue.Front()
	for i := 0; node != nil; i++ {
		input[i] = node.Value.(T)
		node = node.Next()
	}
}
