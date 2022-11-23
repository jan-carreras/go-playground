package double_list

import (
	"fmt"
	"strings"
)

// Implement a nested class DoubleNode for building doubly-linked lists, where
// each Node contains a reference to the item preceding it and the item following
// it in the list (null if there is no such item). Then implement static methods
// for the following tasks:
//
// ✅ - insert at the beginning
// ✅ - insert at the end
// ✅ - remove from the beginning
// ✅ - remove from the end
// ✅ - insert before a given Node
// ✅ - insert after a given Node
// ✅ - remove a given Node.

// Node represents a Node on a double-linked list
type Node[T comparable] struct {
	value    T
	previous *Node[T]
	next     *Node[T]
}

// List represents a double-linked list. Ready to use now `l := new(List[...])`
type List[T comparable] struct {
	first  *Node[T]
	last   *Node[T]
	length int
}

// InsertBeginning inserts an element at the beginning of the list
func (l *List[T]) InsertBeginning(v T) {
	newNode := &Node[T]{value: v}

	if l.first == nil {
		newNode.previous = nil // Clarification
		newNode.next = nil     // Clarification

		l.first = newNode
		l.last = newNode
	} else {
		newNode.previous = nil // Clarification

		l.first.previous = newNode
		newNode.next = l.first
		l.first = newNode
	}

	l.length++
}

func (l *List[T]) RemoveBeginning() {
	if l.isEmpty() {
		return
	} else if l.length == 1 { // One Node
		l.first, l.last = nil, nil
	} else { // Multiple nodes
		second := l.first.next
		second.previous = nil
		l.first = second
	}

	l.length--
}

func (l *List[T]) InsertEnd(v T) {
	// If the list is empty, inserting at the end and begging is the same
	if l.isEmpty() {
		l.InsertBeginning(v)
		return
	}

	newNode := &Node[T]{value: v}
	newNode.next = nil // Clarification

	newNode.previous = l.last
	newNode.next = nil

	l.last.next = newNode

	l.last = newNode

	l.length++
}

func (l *List[T]) RemoveEnd() {
	if l.isEmpty() {
		return
	}

	if l.length == 1 { // One Node: same as RemoveBeginning
		l.RemoveBeginning()
		return
	}

	// Multiple nodes
	previous := l.last.previous

	previous.next = nil
	l.last = previous

	l.length--
}

func (l *List[T]) InsertBefore(n *Node[T], v T) {
	if n.previous == nil { // First node, thus same as inserting in the beginning
		l.InsertBeginning(v)
		return
	}

	newNode := &Node[T]{value: v, previous: n.previous, next: n}

	n.previous.next = newNode
	n.previous = newNode

	l.length++
}

func (l *List[T]) InsertAfter(n *Node[T], v T) {
	if n.next == nil { // Last node, thus same as inserting in the end
		l.InsertEnd(v)
		return
	}

	newNode := &Node[T]{value: v, previous: n, next: n.next}

	n.next.previous = newNode
	n.next = newNode

	l.length++
}

// String returns a visual representation of the list
func (l *List[T]) String() string {
	b := strings.Builder{}
	l.EachN(func(n int, value T) {
		b.WriteString(fmt.Sprintf("%v", value))
		if n != l.length-1 { // Last element on the list
			b.WriteString(" -> ")
		}
	})

	return b.String()
}

// Remove the node from the list.
//
//	Constraint: The node must be in the list!
func (l *List[T]) Remove(n *Node[T]) {
	if l.isEmpty() { // We don't know where this node comes from, but it's not in the list anyway
		return
	}

	if n.previous == nil { // If the node is at the start of the list, we call RemoveBeginning
		if n == l.first { // Defensive programming: this node should be the same as the first
			l.RemoveBeginning()
		}
		return
	}

	if n.next == nil {
		if n == l.last { // Defensive programming: this node should be the same as the last
			l.RemoveEnd() // If the node is at the end of the list, we call RemoveEnd
		}
		return
	}

	prev := n.previous
	next := n.next

	prev.next = next
	next.previous = prev

	l.length--
}

// Each calls fnx for each element passing the value
func (l *List[T]) Each(fnx func(value T)) {
	l.EachN(func(_ int, value T) {
		fnx(value)
	})
}

// EachN calls fnx for each element, passing its position (zero-indexed) and the value
func (l *List[T]) EachN(fnx func(n int, value T)) {
	n := l.first
	for i := 0; n != nil; i++ {
		fnx(i, n.value)
		n = n.next
	}
}

// Length returns the length of the list
func (l *List[T]) Length() int {
	return l.length
}

func (l *List[T]) isEmpty() bool {
	return l.length == 0
}
