package random_queue

import (
	"errors"
	"math/rand"
)

// Random queue. A random queue stores a collection of items and supports the following API
//
// isEmpty
// enqueue
// dequeue
// sample() -> Random element, not removed from the list

var ErrEmpty = errors.New("empty random queue")

type RandomQueue[T comparable] struct {
	List[T]
}

func (r *RandomQueue[T]) Enqueue(v T) {
	r.insertEnd(v)
}

func (r *RandomQueue[T]) Dequeue() (T, error) {
	if r.length == 0 {
		return *new(T), ErrEmpty
	}

	n := r.first
	r.removeBeginning()

	return n.value, nil
}

func (r *RandomQueue[T]) Sample() (T, error) {
	if r.length == 0 {
		return *new(T), ErrEmpty
	}

	randomNode := rand.Intn(r.length)
	n := r.first
	for randomNode > 0 {
		n = n.next
		randomNode--
	}

	return n.value, nil
}
