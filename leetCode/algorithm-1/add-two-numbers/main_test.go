package add_two_numbers

import (
	"fmt"
	"testing"
)

func makeListNode(number int) *ListNode {
	if number == 0 {
		return new(ListNode)
	}

	root := new(ListNode)
	current := root
	for number != 0 {
		current.Val = number % 10
		next := new(ListNode)
		current.Next = next
		current = next

		number /= 10
	}

	return root
}

func TestAddTwoNumbers(t *testing.T) {
	fmt.Println(makeListNode(0))
	fmt.Println(makeListNode(1234))
}
