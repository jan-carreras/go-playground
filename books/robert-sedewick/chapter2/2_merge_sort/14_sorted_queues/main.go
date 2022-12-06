package sorted_queues

import "container/list"

// 2.2.14 Merging sorted queues. Develop a static method that takes two queues of
// sorted items as arguments and returns a queue that results from merging the
// queues into sorted order.

// mergeQueues is as ugly as it gets because I'm not using the "peek" function, only "enqueue/dequeue",
// otherwise the problem is trivial and no different from having two lists and a pointer on each list
// and compare them (see mergeQueuesWithPeek for the trivial response)
func mergeQueues(q1, q2 *list.List) (result *list.List) {
	if q1.Len() == 0 {
		return q2
	}

	if q2.Len() == 0 {
		return q1
	}

	result = list.New()

	e1 := dequeue(q1)
	e2 := dequeue(q2)
	for q1.Len() != 0 && q2.Len() != 0 {
		if e1 < e2 {
			enqueue(result, e1)
			e1 = dequeue(q1)
		} else {
			enqueue(result, e2)
			e2 = dequeue(q2)
		}
	}

	if e1 < e2 {
		enqueue(result, e1)
		enqueue(result, e2)
	} else {
		enqueue(result, e2)
		enqueue(result, e1)
	}

	for q2.Len() != 0 {
		enqueue(result, dequeue(q2))
	}

	for q1.Len() != 0 {
		enqueue(result, dequeue(q1))
	}

	return result
}

func dequeue(q *list.List) int {
	return q.Remove(q.Front()).(int)
}

func enqueue(q *list.List, v int) {
	q.PushBack(v)
}
