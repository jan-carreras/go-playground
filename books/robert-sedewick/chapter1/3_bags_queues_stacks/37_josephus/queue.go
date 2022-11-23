package main

import "errors"

var ErrEmpty = errors.New("empty steque")

type Queue[T comparable] struct {
	List[T]
}

func (s *Queue[T]) Push(v T) {
	s.insertEnd(v)
}

func (s *Queue[T]) Pop() (T, error) {
	if s.isEmpty() {
		return *new(T), ErrEmpty
	}

	n := s.first
	s.removeBeginning()

	return n.value, nil
}
