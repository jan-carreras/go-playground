package main

import (
	"fmt"
	list "leetcode/problems/lists"
)

func main() {
	l1 := list.NewList([]int{1, 3, 4, 12})
	l2 := list.NewList([]int{2, 5, 8, 11})
	l3 := list.NewList([]int{6, 7, 9, 10})

	out := mergeKLists([]*list.ListNode{l1, l2, l3})
	fmt.Println(out.String())
}

func mergeKLists(lists []*list.ListNode) *list.ListNode {
	root := &list.ListNode{}
	current := root

	// Find the minimum node among all elements in the list

	// Iterate until all the lists are empty

	//for i := 0; i < 30; i++ {
	for {
		// If the lists are empty, we have finished
		if len(lists) == 0 {
			return root.Next
		}

		// Search the smallest node from the list
		var newNode *list.ListNode
		for i := 0; i < len(lists); i++ {
			if lists[i] == nil {
				continue
			}
			if newNode == nil || lists[i].Val < newNode.Val {
				newNode = lists[i]
			}
		}
		//fmt.Println(newNode.Val)

		if newNode == nil {
			return root.Next
		}

		// Search from which list we've found this node and move the list forward once
		for i := 0; i < len(lists); i++ {
			if newNode == lists[i] {
				lists[i] = lists[i].Next
				break
			}
		}

		current.Next = newNode
		current = newNode
		newNode = newNode.Next

		// Remove linked lists that are empty
		for i := 0; i < len(lists); i++ {
			if lists[i] == nil {
				lists = append(lists[:i], lists[i+1:]...)
				break // We can only empty one list at a time
			}
		}

		//for _, node := range lists {
		//	fmt.Println("list heads=", node.Val)
		//}

		// We can exit if we've process all the lists

	}
	//panic("probably infinite loop")
}
