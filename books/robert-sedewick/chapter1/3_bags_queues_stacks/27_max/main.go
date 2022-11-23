package main

import (
	"bytes"
	"fmt"
)

// Write a method max() that takes a reference to the first node in a linked list
// as argument and returns the value of the maximum key in the list. Assume that
// all keys are positive integers, and return 0 if the list is empty.

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

func (l *List[T]) Iterate() chan T {
	c := make(chan T, 0)
	if l.length == 0 {
		close(c)
		return c
	}

	n := l.first

	go func() {
		defer close(c)
		for n != nil {
			c <- n.Value
			n = n.Next
		}
	}()

	return c

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

func (l *List[T]) Length() int {
	return l.length
}

func Max(l *List[int]) int {
	max := 0
	for n := range l.Iterate() {
		if n > max {
			max = n
		}
	}

	return max
}
