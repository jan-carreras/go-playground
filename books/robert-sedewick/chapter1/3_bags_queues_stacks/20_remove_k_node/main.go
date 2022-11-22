package _0_remove_k_node

import (
	"bytes"
	"fmt"
)

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type List[T any] struct {
	first, last *Node[T]
	length      int
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

func (l *List[T]) RemoveLast() {
	// Check if we're trying to remove from an empty list
	node := l.first
	if node == nil {
		return
	}

	// If we only have one element
	if l.first == l.last {
		l.length--
		l.first = nil
		l.last = nil
		return
	}

	for node.Next != nil && node.Next.Next != nil {
		node = node.Next
	}

	node.Next = nil
	l.last = node
	l.length--
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
