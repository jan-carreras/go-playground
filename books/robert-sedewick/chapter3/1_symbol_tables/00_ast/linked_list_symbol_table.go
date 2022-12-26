package ast

import (
	"container/list"
	"golang.org/x/exp/constraints"
)

// OrderedLinkedListSymbolTable keeps keys ordered on insert and does a binary search
// to read an element
type OrderedLinkedListSymbolTable[K constraints.Ordered, V any] struct {
	keys   list.List
	values list.List
	size   int
}

func (s *OrderedLinkedListSymbolTable[K, V]) Get(key K) *V {
	index := s.getKeyIndex(key)
	if index == -1 {
		return nil
	}

	return s.getValue(index)
}

func (s *OrderedLinkedListSymbolTable[K, V]) Set(key K, value V) {
	// If the list is empty, we insert the values at the start and call it a day
	if s.isEmpty() {
		s.keys.PushFront(key)
		s.values.PushFront(value)
		return
	}

	// If we find the key, we can update the value and call it aday
	index := s.getKeyIndex(key)
	if index != -1 { // Overwriting the value
		s.getValueNode(index).Value = value

		return
	}

	// If the key does not exist, we need to find where to insert it
	// in our ordered linked list
	node, i := s.keys.Front(), 0
	for ; node.Value.(K) < key; i++ {
		node = node.Next()
	}

	s.keys.InsertBefore(key, node)
	s.values.InsertBefore(value, s.getValueNode(i))
}

func (s *OrderedLinkedListSymbolTable[K, V]) Del(key K) {
	index, node := s.getKeyIndexNode(key)
	if node == nil {
		return // Do nothing, we haven't found it
	}

	s.keys.Remove(node)
	s.values.Remove(s.getValueNode(index))
}

// getKeyIndex returns the index of a given key, -1 if not found
func (s *OrderedLinkedListSymbolTable[K, V]) getKeyIndex(key K) int {
	index, _ := s.getKeyIndexNode(key)
	return index
}

func (s *OrderedLinkedListSymbolTable[K, V]) getKeyIndexNode(key K) (int, *list.Element) {
	node := s.keys.Front()
	for i := 0; node != nil; i++ {
		if node.Value.(K) == key {
			return i, node
		}
		node = node.Next()
	}

	return -1, nil

}

func (s *OrderedLinkedListSymbolTable[K, V]) getValue(index int) *V {
	node := s.getValueNode(index)

	if node == nil {
		return nil
	}

	val := node.Value.(V)
	return &val
}

func (s *OrderedLinkedListSymbolTable[K, V]) getValueNode(index int) *list.Element {
	node := s.values.Front()
	for i := 0; i < index; i++ {
		node = node.Next()
	}

	return node
}

func (s *OrderedLinkedListSymbolTable[K, V]) isFull() bool {
	return s.keys.Len() == s.size
}

func (s *OrderedLinkedListSymbolTable[K, V]) isEmpty() bool {
	return s.keys.Len() == 0
}
