package linkedlist

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.

type List struct {
	root *Node
}

type Node struct {
	Value interface{}
	list  *List
	next  *Node
	prev  *Node
}

func NewList(args ...interface{}) *List {
	root := &Node{}
	root.next = root
	root.prev = root

	l := &List{root: root}
	root.list = l

	for _, v := range args {
		l.Push(v)
	}

	return l
}

func (n *Node) isRoot() bool {
	return n == n.list.root
}

func (n *Node) Next() *Node {
	if n.next.isRoot() {
		return nil
	}

	return n.next
}

func (n *Node) Prev() *Node {
	if n.prev.isRoot() {
		return nil
	}

	return n.prev
}

// Unshift inserts values at the front
func (l *List) Unshift(v interface{}) {
	n := &Node{Value: v, list: l}

	n.next = l.root.next
	n.prev = l.root

	n.next.prev = n
	n.prev.next = n
}

// Push inserts value at the back
func (l *List) Push(v interface{}) {
	n := &Node{Value: v, list: l}

	n.next = l.root
	n.prev = l.root.prev

	n.next.prev = n
	n.prev.next = n
}

// Shift removes values at the front
func (l *List) Shift() (interface{}, error) {
	first := l.root.next

	first.next.prev = l.root
	first.prev.next = first.next

	val := first.Value
	first = nil

	return val, nil
}

// Pop removes value at the back
func (l *List) Pop() (interface{}, error) {
	last := l.root.prev

	last.prev.next = l.root
	last.next.prev = last.prev

	val := last.Value
	last = nil

	return val, nil
}

func (l *List) Reverse() {

	n := l.First()
	for n != nil {
		nextNode := n.Next()
		n.prev, n.next = n.next, n.prev // Swap the values!

		n = nextNode // Transverse the list
	}

	l.root.next, l.root.prev = l.root.prev, l.root.next
}

func (l *List) First() *Node {
	return l.root.Next()
}

func (l *List) Last() *Node {
	return l.root.Prev()
}
