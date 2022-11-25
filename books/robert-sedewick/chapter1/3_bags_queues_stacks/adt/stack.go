package adt

import (
	"errors"
)

var ErrEmptyStack = errors.New("empty stack")

type Stack struct {
	list List
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(v any) {
	s.list.PushFront(v)
}

func (s *Stack) Pop() (any, error) {
	if s.list.Len() == 0 {
		return nil, ErrEmptyStack
	}
	n := s.list.Front()
	s.list.Remove(n)

	return n.Value, nil
}

func (s *Stack) Len() int {
	return s.list.Len()
}

func (s *Stack) String() string {
	return String(s.list)
}
