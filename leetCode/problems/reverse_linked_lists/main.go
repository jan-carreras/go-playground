package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func main() {
	root := Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3, Next: &Node{Value: 4}}}}
	print(&root)
	inversed := returnLast(&root)
	print(inversed)
}

func returnLast(head *Node) *Node {
	current := head
	var prev, next *Node
	for current != nil {
		next = current.Next
		current.Next = prev

		prev = current
		current = next
	}

	return prev
}

func print(root *Node) {
	var count int
	n := root
	for n != nil {
		fmt.Println(n.Value)
		n = n.Next
		count++
		if count > 10 {
			panic("meh")
		}
	}
	fmt.Println()
}
