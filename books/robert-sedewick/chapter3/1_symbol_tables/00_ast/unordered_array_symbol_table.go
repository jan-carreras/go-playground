package ast

import (
	"golang.org/x/exp/constraints"
)

// UnorderedSymbolTable keeps keys ordered on insert and does a binary search
// to read an element
type UnorderedSymbolTable[K constraints.Ordered, V any] struct {
	keys   []K
	values []V
	size   int
}

func (s *UnorderedSymbolTable[K, V]) Get(key K) *V {
	i := s.getIndex(key)
	if i == -1 {
		return nil
	}

	return &s.values[i]
}

func (s *UnorderedSymbolTable[K, V]) Set(key K, value V) {
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

	s.add(key, value)
}

func (s *UnorderedSymbolTable[K, V]) Del(key K) {
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
func (s *UnorderedSymbolTable[K, V]) getIndex(key K) int {
	for i := 0; i < s.size; i++ {
		if s.keys[i] == key {
			return i
		}
	}

	return -1
}

func (s *UnorderedSymbolTable[K, V]) add(key K, value V) {
	s.keys[s.size], s.values[s.size] = key, value
	s.size++
}

func (s *UnorderedSymbolTable[K, V]) isFull() bool {
	return len(s.keys) == s.size
}

func (s *UnorderedSymbolTable[K, V]) isEmpty() bool {
	return len(s.keys) == 0
}

func (s *UnorderedSymbolTable[K, V]) resize(n int) {
	keys := make([]K, n)
	copy(keys, s.keys)
	s.keys = keys

	values := make([]V, n)
	copy(values, s.values)
	s.values = values
}
