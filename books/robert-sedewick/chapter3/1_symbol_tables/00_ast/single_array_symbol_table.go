package ast

import "golang.org/x/exp/constraints"

type keyValue[K constraints.Ordered, V any] struct {
	key   K
	value V
}

// SingleArraySymbolTable keeps keys ordered on insert and does a binary search
// to read an element
type SingleArraySymbolTable[K constraints.Ordered, V any] struct {
	kv   []keyValue[K, V]
	size int
}

func (s *SingleArraySymbolTable[K, V]) Get(key K) *V {
	kv := s.getKeyValue(key)
	if kv == nil {
		return nil
	}

	return &kv.value
}

func (s *SingleArraySymbolTable[K, V]) Set(key K, value V) {
	if s.isEmpty() {
		s.resize(1)
	}

	kv := s.getKeyValue(key)
	if kv != nil {
		kv.value = value
		return
	}

	if s.isFull() {
		s.resize(len(s.kv) * 2)
	}

	s.addInPosition(s.searchInsertPosition(key), key, value)
}

func (s *SingleArraySymbolTable[K, V]) Del(key K) {
	idx := s.getIndex(key)
	if idx == -1 { // Nothing to remove
		return
	}

	// Move all elements one position to the right
	for i := idx; i < s.size-1; i++ {
		s.kv[i].key, s.kv[i].value = s.kv[i+1].key, s.kv[i+1].value
	}
	s.size--

	if float64(s.size)/float64(len(s.kv)) < 0.25 {
		s.resize(len(s.kv) / 2)
	}
}

func (s *SingleArraySymbolTable[K, V]) getKeyValue(key K) *keyValue[K, V] {
	index := s.getIndex(key)
	if index == -1 {
		return nil
	}

	return &s.kv[index]
}

// getIndex returns the index of a given key, -1 if not found
func (s *SingleArraySymbolTable[K, V]) getIndex(key K) (mid int) {
	lo, hi := 0, s.size-1
	for hi >= lo {
		mid = lo + ((hi - lo) / 2)
		if s.kv[mid].key > key {
			hi = mid - 1
		} else if s.kv[mid].key < key {
			lo = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

func (s *SingleArraySymbolTable[K, V]) searchInsertPosition(key K) int {
	lo, hi := 0, s.size-1
	for hi >= lo {
		mid := lo + ((hi - lo) / 2)
		if s.kv[mid].key > key {
			hi = mid - 1
		} else if s.kv[mid].key < key {
			lo = mid + 1
		} else {
			return mid
		}
	}

	return lo
}

func (s *SingleArraySymbolTable[K, V]) addInPosition(i int, key K, value V) {
	// Move all elements one position to the left
	for j := s.size; j > i; j-- {
		s.kv[j].key, s.kv[j].value = s.kv[j-1].key, s.kv[j-1].value
	}
	s.kv[i].key, s.kv[i].value = key, value
	s.size++
}

func (s *SingleArraySymbolTable[K, V]) Keys() []K {
	r := make([]K, 0, s.size)
	for i := 0; i < s.size; i++ {
		r = append(r, s.kv[i].key)
	}

	return r
}

func (s *SingleArraySymbolTable[K, V]) isFull() bool {
	return len(s.kv) == s.size
}

func (s *SingleArraySymbolTable[K, V]) isEmpty() bool {
	return len(s.kv) == 0
}

func (s *SingleArraySymbolTable[K, V]) resize(n int) {
	keys := make([]keyValue[K, V], n)
	copy(keys, s.kv)
	s.kv = keys
}
