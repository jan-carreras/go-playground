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

	var newNode *list.ListNode
	newNodeIndex := 0
	for {
		newNode, newNodeIndex = nil, 0

		for i := 0; i < len(lists); i++ {
			if lists[i] == nil {
				continue
			}
			if newNode == nil || lists[i].Val < newNode.Val {
				newNode = lists[i]
				newNodeIndex = i
			}
		}

		if newNode == nil {
			return root.Next
		}

		// Search from which list we've found this node and move the list forward once
		lists[newNodeIndex] = lists[newNodeIndex].Next

		current.Next = newNode
		current = newNode
		newNode = newNode.Next
	}
}
