package main

import (
	"bytes"
	"fmt"
)

// 1.3.24 Write a method remove() that takes a linked list and a string key as
// arguments and removes all the nodes in the list that have key as its item
// field.

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type List[T comparable] struct {
	first, last *Node[T]
	length      int
}

func (l *List[T]) String() string {
	b := bytes.Buffer{}
	n := l.first
	for n != nil {
		b.WriteString(fmt.Sprintf("%v", n.Value))
		if n.Next != nil {
			b.WriteString(" -> ")
		}
		n = n.Next
	}

	return b.String()
}

// RemoveAll this method should not exist! It exposes the internals of the linked list!
// TODO: The problem with this implementation is that we need to iterate the entire List
//
//	for each deletion, instead of doing it in-place. If we want to remove it InPlace
//	we could be using a RemoveAfter function to achieve the same goal
func (l *List[T]) RemoveAll(key T) {
	for l.first != nil && l.first.Value == key {
		l.RemoveK(0)
	}
	if l.first == nil {
		return
	}

	n := l.first
	index := 0
	for ; n != nil; index++ {
		if n.Value == key {
			l.RemoveK(index)
			index-- // One less item
		}
		n = n.Next
	}
}

// RemoveK removed the Kth element. k is 0 index
func (l *List[T]) RemoveK(k int) {
	// We're trying to remove an element out of bounds. NoOp.
	if k > l.length {
		return
	}

	// Removing the first element is an edge case
	if k == 0 {
		l.first = l.first.Next
		l.length--

		// More edge cases: If the new length is 0, both l.first and l.last must point at
		// the same place
		if l.length == 0 {
			l.last = l.first
		}

		return
	}

	// Locate the node BEFORE k
	n := l.first
	for i := 0; i < k-1; i++ {
		n = n.Next
		if n == nil {
			// This should not be possible, really. We've checked if we're in bounds, but
			// you never are sure enough
			return
		}
	}

	// Remove the actual node
	n.Next = n.Next.Next
	l.length--

	// If we have removed the last node, we need to point l.last to that Node
	if n.Next == nil {
		l.last = n
	}

}

func (l *List[T]) Enqueue(val T) {
	newNode := &Node[T]{
		Value: val,
	}
	if l.length == 0 {
		l.first = newNode
		l.last = newNode
	} else {
		l.last.Next = newNode
		l.last = newNode
	}

	l.length++
}

func (l *List[T]) Dequeue() T {
	if l.length == 0 {
		panic("queue is empty") // Could be a classic error
	}

	res := l.first
	l.first = l.first.Next

	l.length--
	return res.Value
}
