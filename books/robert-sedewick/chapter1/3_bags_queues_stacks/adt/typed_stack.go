package adt

type TypeStack[T any] struct {
	list List
}

func NewTypeStack[T any]() *TypeStack[T] {
	return &TypeStack[T]{}
}

func (s *TypeStack[T]) Push(v T) {
	s.list.PushFront(v)
}

func (s *TypeStack[T]) Pop() (T, error) {
	if s.list.Len() == 0 {
		return *new(T), ErrEmptyStack
	}
	n := s.list.Front()
	s.list.Remove(n)

	return n.Value.(T), nil
}

// SPop is the same as Pop, but silent (never errors even when empty)
func (s *TypeStack[T]) SPop() T {
	v, err := s.Pop()
	if err != nil {
		return *new(T)
	}

	return v
}

func (s *TypeStack[T]) Len() int {
	return s.list.Len()
}

func (s *TypeStack[T]) String() string {
	return String(s.list)
}
