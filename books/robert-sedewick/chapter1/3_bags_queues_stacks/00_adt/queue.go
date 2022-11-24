package adt

import (
	"container/list"
	"errors"
	format "exercises/books/robert-sedewick/chapter1/3_bags_queues_stacks/00_format"
)

var ErrEmptyQueue = errors.New("empty queue")

type Queue struct {
	list list.List
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Len() int {
	return q.list.Len()
}

func (q *Queue) Enqueue(v any) {
	q.list.PushBack(v)
}

func (q *Queue) Dequeue() (any, error) {
	if q.list.Len() == 0 {
		return nil, ErrEmptyQueue
	}

	n := q.list.Front()
	q.list.Remove(n)

	return n.Value, nil
}

func (q *Queue) String() string {
	return format.List(q.list)
}
