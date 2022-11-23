package main

import (
	"bytes"
	"fmt"
)

// 1.3.21 Write a method find() that takes a linked list and a string key as
// arguments and returns true if some node in the list has key as its item field,
// false otherwise.

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type List[T comparable] struct {
	first, last *Node[T]
	length      int
}

func (l *List[T]) Find(search T) bool {
	n := l.first
	for n != nil {
		if n.Value == search {
			return true
		}
		n = n.Next
	}

	return false
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
