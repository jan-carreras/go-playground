package main

import (
	"errors"
	adt "exercises/books/robert-sedewick/chapter1/3_bags_queues_stacks/00_adt"
)

// Copy a queue. Create a new constructor so that:
//
// Queue<Item> r = new Queue<Item>(q);
//
// makes `r` a reference to a new and independent copy of the queue `q`. You
// should be able to push and pop from either q or r without influencing the
// other.
//
// Hint: Delete all the elements from q and add these elements to both q and r.

var ErrEmpty = errors.New("empty queue")

type Queue struct {
	list adt.List
}

func NewQueue(oldQueue *Queue) *Queue {
	newQueue := &Queue{}

	for i := 0; i < oldQueue.Len(); i++ {
		v, err := oldQueue.Dequeue()
		if err != nil {
			panic(err) // The queue has been modified while copying
		}

		oldQueue.Enqueue(v)
		newQueue.Enqueue(v)
	}

	return newQueue
}

func (q *Queue) Len() int {
	return q.list.Len()
}

func (q *Queue) Enqueue(v any) {
	q.list.PushBack(v)
}

func (q *Queue) Dequeue() (any, error) {
	if q.list.Len() == 0 {
		return nil, ErrEmpty
	}

	n := q.list.Front()
	q.list.Remove(n)

	return n.Value, nil
}

func (q *Queue) String() string {
	return adt.String(q.list)
}
