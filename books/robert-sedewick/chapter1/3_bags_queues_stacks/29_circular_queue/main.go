package main

import (
	"errors"
	"fmt"
	"strings"
)

var ErrEmptyQueue = errors.New("queue is empty")

// 1.3.29 Write a Queue implementation that uses a circular linked list, which is
// the same as a linked list except that no links are null and the value of
// last.next is first whenever the list is not empty. Keep only one Node
// instance variable (last).

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

type CircularQueue[T comparable] struct {
	length int
	last   *Node[T]
}

func (q *CircularQueue[T]) String() string {
	if q.length == 0 {
		return ""
	}

	b := strings.Builder{}
	q.EachN(func(i int, value T) {
		b.WriteString(fmt.Sprintf("%v", value))
		if i != q.length-1 { // Not last element
			b.WriteString(" -> ")
		}
	})

	return b.String()
}

func (q *CircularQueue[T]) Enqueue(value T) {
	newNode := &Node[T]{value: value}

	// When the list is empty
	if q.last == nil {
		q.last = newNode
		q.last.next = q.last // Point to itself
	} else { // The circular list is non-empty
		newNode.next = q.last.next
		q.last.next = newNode
		q.last = newNode
	}

	q.length++
}

func (q *CircularQueue[T]) Dequeue() (rsp T, err error) {
	if q.last == nil {
		return *new(T), ErrEmptyQueue
	}

	// Get the first element on the list
	rsp = q.last.next.value

	q.last.next = q.last.next.next // Remove First element

	q.length--

	if q.length == 0 {
		q.last = nil
	}

	return rsp, nil
}

func (q *CircularQueue[T]) Length() int {
	return q.length
}

// EachN iterates all the elements on the Queue, calling fnx with the zero-index position and the value
func (q *CircularQueue[T]) EachN(fnx func(i int, value T)) {
	n := q.last.next
	for i := 0; i < q.length; i++ {
		fnx(i, n.value)
		n = n.next
	}
}

// Each iterates all the elements on the Queue, calling fnx value
func (q *CircularQueue[T]) Each(fnx func(value T)) {
	q.EachN(func(_ int, value T) {
		fnx(value)
	})
}
