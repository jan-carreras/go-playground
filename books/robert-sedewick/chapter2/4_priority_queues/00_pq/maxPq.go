package pq

type Value[T any] struct {
	value    T
	priority int
}

func NewValue[T any](value T, priority int) Value[T] {
	return Value[T]{value: value, priority: priority}
}

func (s Value[T]) Value() T {
	return s.value
}

func (s Value[_]) Priority() int {
	return s.priority
}

// MaxPQ stores the N minimum inserted values
type MaxPQ[T any] struct {
	size int
	heap []Value[T]
}

func NewMaxPQ[T any](maxN int) *MaxPQ[T] {
	return &MaxPQ[T]{
		size: 0,
		heap: make([]Value[T], maxN+1), // The heap[0] is unused
	}
}

func (m *MaxPQ[T]) isFull() bool {
	return m.size == len(m.heap)-1
}

func (m *MaxPQ[KP]) IsEmpty() bool {
	return m.Size() == 0
}

func (m *MaxPQ[T]) resize(newSize int) {
	aux := make([]Value[T], newSize)
	copy(aux, m.heap)
	m.heap = aux
}

func (m *MaxPQ[T]) Insert(keyPriority Value[T]) {
	if m.isFull() {
		m.resize(len(m.heap) * 2)
	}

	m.size++
	m.heap[m.size] = keyPriority
	m.swim(m.size)
}

func (m *MaxPQ[T]) DeleteMax() (Value[T], bool) {
	if m.size == 0 {
		return *(new(Value[T])), false // Priority Queue underflow
	}

	key := m.heap[1]                // Max value. m.heap[0] is empty/ignored
	m.exchange(1, m.size)           // Exchange with last item
	m.heap[m.size] = *new(Value[T]) // Help garbage collection to clean up
	m.size--
	m.sink(1) // Restore heap property

	if m.size > 0 && m.size == (len(m.heap)-1)/4 {
		m.resize(len(m.heap) / 2)
	}
	return key, true
}

func (m *MaxPQ[_]) Size() int {
	return m.size
}

func (m *MaxPQ[_]) less(i, j int) bool {
	return m.heap[i].Priority() < m.heap[j].Priority()
}

func (m *MaxPQ[_]) exchange(i, j int) {
	m.heap[i], m.heap[j] = m.heap[j], m.heap[i]
}

func (m *MaxPQ[_]) swim(k int) {
	for k > 1 && m.less(k/2, k) {
		m.exchange(k/2, k)
		k /= 2
	}
}

func (m *MaxPQ[_]) sink(k int) {
	for 2*k <= m.size {
		j := 2 * k
		if j < m.size && m.less(j, j+1) {
			j++
		}
		if !m.less(k, j) {
			break
		}
		m.exchange(k, j)
		k = j
	}
}
