package deque

import "errors"

// 1.3.33 Deque. A double-ended queue or deque (pronounced “deck”) is like a
// stack or a queue but supports adding and removing items at both ends. A deque
// stores a collection of items and supports the following API:
//
// ✅ isEmpty() bool
// ✅ size() int
// ✅ pushLeft(item)
// ✅ pushRight(item)
// ✅ popLeft() Item
// ✅ popRight() Item

var ErrEmpty = errors.New("empty deque")

type Steque[T comparable] struct {
	List[T]
}

func (s *Steque[T]) IsEmpty() bool {
	return s.isEmpty()
}

func (s *Steque[T]) PushLeft(v T) {
	s.insertBeginning(v)
}

func (s *Steque[T]) PushRight(v T) {
	s.insertEnd(v)
}

func (s *Steque[T]) PopLeft() (T, error) {
	if s.isEmpty() {
		return *new(T), ErrEmpty
	}

	n := s.first
	s.removeBeginning()

	return n.value, nil
}

func (s *Steque[T]) PopRight() (T, error) {
	if s.isEmpty() {
		return *new(T), ErrEmpty
	}

	n := s.last
	s.removeEnd()

	return n.value, nil
}
