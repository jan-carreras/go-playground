package list

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(input []int) *ListNode {
	root := &ListNode{}
	current := root

	for _, i := range input {
		newL := &ListNode{Val: i}
		current.Next = newL
		current = newL
	}

	return root.Next
}

func (l *ListNode) String() string {
	b := strings.Builder{}

	var count int

	n := l
	for n != nil {
		b.WriteString(fmt.Sprintf("%d\n", n.Val))
		n = n.Next

		if count > 100 {
			b.WriteString("[PANIC] YOU ARE PROBABLY PRINTING IN A LOOP")
			return b.String()
		}
		count++
	}

	return b.String()
}
