package ast

import (
	"golang.org/x/exp/constraints"
)

// ArraySymbolTable keeps keys ordered on insert and does a binary search
// to read an element
type ArraySymbolTable[K constraints.Ordered, V any] struct {
	keys   []K
	values []V
	size   int
}

func (s *ArraySymbolTable[K, V]) Get(key K) *V {
	i := s.getIndex(key)
	if i == -1 {
		return nil
	}

	return &s.values[i]
}

func (s *ArraySymbolTable[K, V]) Set(key K, value V) {
	if s.isEmpty() {
		s.resize(1)
	}

	i := s.getIndex(key)
	if i != -1 { // Overwriting the value
		s.values[i] = value
		return
	}

	if s.isFull() {
		s.resize(len(s.keys) * 2)
	}

	s.addInPosition(s.searchInsertPosition(key), key, value)
}

func (s *ArraySymbolTable[K, V]) Del(key K) {
	idx := s.getIndex(key)
	if idx == -1 { // Nothing to remove
		return
	}

	// Move all elements one position to the right
	for i := idx; i < s.size-1; i++ {
		s.keys[i], s.values[i] = s.keys[i+1], s.values[i+1]
	}
	s.size--

	if float64(s.size)/float64(len(s.keys)) < 0.25 {
		s.resize(len(s.keys) / 2)
	}
}

// getIndex returns the index of a given key, -1 if not found
func (s *ArraySymbolTable[K, V]) getIndex(key K) (mid int) {
	lo, hi := 0, s.size-1
	for hi >= lo {
		mid = lo + ((hi - lo) / 2)
		if s.keys[mid] > key {
			hi = mid - 1
		} else if s.keys[mid] < key {
			lo = mid + 1
		} else {
			return mid
		}
	}

	return -1

}

func (s *ArraySymbolTable[K, V]) searchInsertPosition(key K) int {
	lo, hi := 0, s.size-1
	for hi >= lo {
		mid := lo + ((hi - lo) / 2)
		if s.keys[mid] > key {
			hi = mid - 1
		} else if s.keys[mid] < key {
			lo = mid + 1
		} else {
			return mid
		}
	}

	return lo
}

func (s *ArraySymbolTable[K, V]) addInPosition(i int, key K, value V) {
	// Move all elements one position to the left
	for j := s.size; j > i; j-- {
		s.keys[j], s.values[j] = s.keys[j-1], s.values[j-1]
	}
	s.keys[i], s.values[i] = key, value
	s.size++
}

func (s *ArraySymbolTable[K, V]) isFull() bool {
	return len(s.keys) == s.size
}

func (s *ArraySymbolTable[K, V]) isEmpty() bool {
	return len(s.keys) == 0
}

func (s *ArraySymbolTable[K, V]) resize(n int) {
	keys := make([]K, n)
	copy(keys, s.keys)
	s.keys = keys

	values := make([]V, n)
	copy(values, s.values)
	s.values = values
}
