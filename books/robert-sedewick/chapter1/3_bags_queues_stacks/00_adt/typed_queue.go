package adt

type TypedQueue[T any] struct {
	list List
}

func NewTypedQueue[T any]() *TypedQueue[T] {
	return &TypedQueue[T]{}
}

func (q *TypedQueue[T]) Len() int {
	return q.list.Len()
}

func (q *TypedQueue[T]) Enqueue(v any) {
	q.list.PushBack(v)
}

func (q *TypedQueue[T]) SDequeue() any {
	v, err := q.Dequeue()
	if err != nil {
		return *new(T)
	}
	return v
}

func (q *TypedQueue[T]) Dequeue() (any, error) {
	if q.list.Len() == 0 {
		return nil, ErrEmptyQueue
	}

	n := q.list.Front()
	q.list.Remove(n)

	return n.Value, nil
}

func (q *TypedQueue[T]) String() string {
	return String(q.list)
}

// RemoveN remove the nth element in the stack
func (q *TypedQueue[T]) RemoveN(n int) {
	if n > q.Len() {
		return // No Op. Nothing to remove
	}

	node := q.list.Front()
	for i := 0; i < n; i++ {
		node = node.Next()
	}
	q.list.Remove(node)
}

// EachN calls fnx for each value on the list. Stops if fnx returns false
func (q *TypedQueue[T]) EachN(fnx func(n int, v any) bool) {
	n := q.list.Front()
	for i := 0; *n != q.list.root; i++ {
		if !fnx(i, n.Value) {
			return
		}
		n = n.next
	}
}
