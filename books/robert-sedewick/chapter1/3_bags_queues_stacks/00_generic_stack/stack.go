package stack

import (
	"bytes"
	"fmt"
)

// Possible improvements: do not panic, return a custom error instead

type Stack[T any] struct {
	first  *Node[T]
	length int
}

func (s *Stack[T]) Length() int {
	return s.length
}

func (s *Stack[T]) Push(val T) {
	n := &Node[T]{value: val, next: s.first}
	s.first = n
	s.length++
}

func (s *Stack[T]) Pop() T {
	if s.length == 0 {
		panic("empty!!")
	}

	n := s.first
	s.first = n.next

	s.length--
	return n.value
}

func (s *Stack[T]) String() string {

	buf := bytes.Buffer{}

	n := s.first
	for n != nil {
		buf.WriteString(fmt.Sprintf("--> %v ", n.value))
		n = n.next
	}

	return buf.String()
}
