package steque

import "errors"

// 1.3.32 Steque. A stack-ended queue or steque is a data type that supports
//
// ✅ push
// ✅ pop
// ✅ enqueue
//
// Articulate an API for this ADT. Develop a linked-list-based implementation.

var ErrEmpty = errors.New("empty steque")

type Steque[T comparable] struct {
	List[T]
}

func (s *Steque[T]) Push(v T) {
	s.insertBeginning(v)
}

func (s *Steque[T]) Pop() (T, error) {
	if s.isEmpty() {
		return *new(T), ErrEmpty
	}

	n := s.first
	s.removeBeginning()

	return n.value, nil
}

func (s *Steque[T]) Enqueue(v T) {
	s.insertEnd(v)
}
