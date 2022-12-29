package lru

import (
	"container/list"
	"golang.org/x/exp/constraints"
)

// 3.1.22 Self-organizing search. A self-organizing search algorithm is one that
// rearranges items to make those that are accessed frequently likely to be
// found early in the search. Modify your search implementation for Exercise
// 3.1.2 to perform the following action on every search hit: move the key-value
// pair found to the beginning of the list, moving all pairs between the
// beginning of the list and the vacated position to the right one position.
// This procedure is called the move-to-front heuristic.

type LRULinkedListSymbolTable[K constraints.Ordered, V any] struct {
	keys   list.List
	values list.List
	size   int
}

func (s *LRULinkedListSymbolTable[K, V]) Get(key K) *V {
	index, keyNode := s.getKeyIndexNode(key)
	if index == -1 {
		return nil
	}

	valueNode := s.getValueNode(index)

	s.keys.MoveToFront(keyNode)
	s.values.MoveToFront(valueNode)

	value := valueNode.Value.(V)
	return &value
}

func (s *LRULinkedListSymbolTable[K, V]) Set(key K, value V) {
	// If the list is empty, we insert the values at the start and call it a day
	if s.isEmpty() {
		s.keys.PushFront(key)
		s.values.PushFront(value)
		return
	}

	// If we find the key, we can update the value and call it day
	index, keyNode := s.getKeyIndexNode(key)
	if index != -1 { // Overwriting the value
		valueNode := s.getValueNode(index)
		valueNode.Value = value

		// Moving the key-value at the start of the list
		s.keys.MoveToFront(keyNode)
		s.values.MoveToFront(valueNode)

		return
	}

	s.keys.PushFront(key)
	s.values.PushFront(value)
}

func (s *LRULinkedListSymbolTable[K, V]) Del(key K) {
	index, node := s.getKeyIndexNode(key)
	if node == nil {
		return // Do nothing, we haven't found it
	}

	s.keys.Remove(node)
	s.values.Remove(s.getValueNode(index))
}

// getKeyIndex returns the index of a given key, -1 if not found
func (s *LRULinkedListSymbolTable[K, V]) getKeyIndex(key K) int {
	index, _ := s.getKeyIndexNode(key)
	return index
}

func (s *LRULinkedListSymbolTable[K, V]) getKeyIndexNode(key K) (int, *list.Element) {
	node := s.keys.Front()
	for i := 0; node != nil; i++ {
		if node.Value.(K) == key {
			return i, node
		}
		node = node.Next()
	}

	return -1, nil

}

func (s *LRULinkedListSymbolTable[K, V]) getValue(index int) *V {
	node := s.getValueNode(index)

	if node == nil {
		return nil
	}

	val := node.Value.(V)
	return &val
}

func (s *LRULinkedListSymbolTable[K, V]) getValueNode(index int) *list.Element {
	node := s.values.Front()
	for i := 0; i < index; i++ {
		node = node.Next()
	}

	return node
}

func (s *LRULinkedListSymbolTable[K, V]) isFull() bool {
	return s.keys.Len() == s.size
}

func (s *LRULinkedListSymbolTable[K, V]) isEmpty() bool {
	return s.keys.Len() == 0
}
