package main

import (
	"container/list"
	"fmt"
	"strings"
)

type MoveToFront struct {
	list list.List
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
	b := strings.Builder{}
	n := m.list.Front()
	for n != nil {
		b.WriteString(fmt.Sprintf("%v -> ", n.Value))
		n = n.Next()
	}

	return b.String()
}
