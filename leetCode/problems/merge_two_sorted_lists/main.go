package main

import (
	"fmt"
	list "leetcode/problems/lists"
)

func main() {
	l1 := list.NewList([]int{1, 2, 4})
	l2 := list.NewList([]int{1, 3, 4})

	out := mergeTwoLists(l1, l2)
	fmt.Println(out)
}

func mergeTwoLists(list1 *list.ListNode, list2 *list.ListNode) *list.ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	root := &list.ListNode{}
	current := root

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
			current = current.Next
		} else {
			current.Next = list2
			list2 = list2.Next
			current = current.Next
		}
	}

	for list1 != nil {
		current.Next = list1
		list1 = list1.Next
		current = current.Next
	}

	for list2 != nil {
		current.Next = list2
		list2 = list2.Next
		current = current.Next
	}

	return root.Next
}
