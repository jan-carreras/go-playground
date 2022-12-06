package sorted_queues

import "container/list"

var Merge = mergeQueuesWithPeek

func mergeQueuesWithPeek(q1, q2 *list.List) (result *list.List) {
	result = list.New()

	for q1.Len() != 0 && q2.Len() != 0 {
		if peek(q1) < peek(q2) {
			enqueue(result, dequeue(q1))
		} else {
			enqueue(result, dequeue(q2))
		}
	}

	for q2.Len() != 0 {
		enqueue(result, dequeue(q2))
	}

	for q1.Len() != 0 {
		enqueue(result, dequeue(q1))
	}

	return result
}

func peek(q *list.List) int {
	return q.Front().Value.(int)
}
