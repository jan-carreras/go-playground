package main

import adt "github.com/jan-carreras/go-playground/books/robert-sedewick/chapter1/3_bags_queues_stacks/adt"

// 1.3.40 Move-to-front. Read in a sequence of characters from standard input and
// maintain the characters in a linked list with no duplicates. When you read in
// a previously unseen character, insert it at the front of the list. When you
// read in a duplicate character, delete it from the list and reinsert it at the
// beginning. Name your program MoveToFront: it implements the well-known
// move-to-front strategy, which is useful for caching, data compression, and
// many other applications where items that have been recently accessed are more
// likely to be reaccessed

type MoveToFront struct {
	list adt.List
}

func (m *MoveToFront) Add(v any) {
	m.remove(v)
	m.list.PushFront(v)
}

func (m *MoveToFront) remove(v any) {
	n := m.list.Front()
	for n != nil {
		if n.Value == v {
			m.list.Remove(n)
			break
		}
		n = n.Next()
	}
}

func (m *MoveToFront) String() string {
	return adt.String(m.list)
}
