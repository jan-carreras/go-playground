package main

import (
	list "leetcode/problems/lists"
)

func main() {

}

func removeNthFromEnd(head *list.ListNode, n int) *list.ListNode {
	// Nothing to remove, just return original
	if n == 0 {
		return head
	}

	fast, slow := head, head
	for fast != nil {
		fast = fast.Next
		if n >= 0 {
			n--
			continue
		}
		slow = slow.Next
	}

	if n > 0 { // We try to delete an element that it's not in the list. NoOP.
		return head
	} else if n == 0 { // Try to remove the first element on the list
		return head.Next
	} else { // Remove the specific element
		slow.Next = slow.Next.Next
	}

	return head

}
