package main

import (
	"bytes"
	"fmt"
)

// 1.3.24 Write a method insertAfter() that takes two linked-list Node arguments
// and inserts the second after the first on its list (and does nothing if either
// argument is null).

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

// InsertAfter this method should not exist! It exposes the internals of the linked list!
func (l *List[T]) InsertAfter(node *Node[T], newNode *Node[T]) {
	newNode.Next = node.Next
	node.Next = newNode

	if newNode.Next == nil {
		l.last = newNode
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
