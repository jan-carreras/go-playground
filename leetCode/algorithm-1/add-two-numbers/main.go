package add_two_numbers

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() (str string) {
	next := l
	for next != nil {
		str += fmt.Sprintf("-> [%d] ", next.Val)
		next = next.Next
	}

	return str
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	return nil
}
