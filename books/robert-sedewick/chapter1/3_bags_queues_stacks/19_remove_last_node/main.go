package main

import (
	"bytes"
	"fmt"
)

// 1.3.19 Give a code fragment that removes the last node in a linked list whose
// first node is `first`.

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type List[T any] struct {
	first, last *Node[T]
	length      int
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
