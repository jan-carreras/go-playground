package random_order

import "math/rand"

type RandomBag[T comparable] struct {
	List[T]
}

func (r *RandomBag[T]) IsEmpty() bool {
	return r.isEmpty()
}

func (r *RandomBag[T]) Size() int {
	return r.length
}

func (r *RandomBag[T]) Add(v T) {
	if r.length == 0 {
		r.insertBeginning(v)
		return
	}

	randomPosition := rand.Intn(r.length)
	n := r.first
	for randomPosition > 0 {
		n = n.next
		randomPosition--
	}

	if rand.Intn(2)%2 == 0 {
		r.insertBefore(n, v)
	} else {
		r.insertAfter(n, v)
	}
}
