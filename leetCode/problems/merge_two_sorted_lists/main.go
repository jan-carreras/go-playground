package main

import "fmt"

func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	out := mergeTwoLists(l1, l2)
	print(out)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	root := &ListNode{}
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

func print(root *ListNode) {
	var count int
	n := root
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
		count++
		if count > 10 {
			panic("meh")
		}
	}
	fmt.Println()
}
