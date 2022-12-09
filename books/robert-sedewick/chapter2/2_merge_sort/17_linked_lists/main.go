package linked_lists

import (
	"container/list"
	"fmt"
)

// 2.2.17 Linked-list sort.
//
// Implement a natural mergesort for linked lists. (This is the method of choice
// for sorting linked lists because it uses no extra space and is guaranteed to
// be linearithmic.)

// Note: Using *list.List for convenience, but I'm only going to use the Next object
// because the exercise doesn't mention doubly-linked lists.

// TODO: List elements must be integers. This should be fixed by using types and a proper wrapper for
// a linked list implementation

func MergeSort(input *list.List) {
	// search for 1 ascending group. search for a second ascending group
	//	a scending group is described before no element it bigger than the previous one
	// call the merge function
	//   iterate from mid to hi, and put the values in the correct places on lo->mid
	//   keep a pointer on lo, and compare it with the increasing mid. If the value in mid
	//   is smaller than the current in lo, insert before
	//
	// Unit test every part of the code
	if input == nil || input.Front() == nil || input.Front().Next() == nil {
		return
	}

SORTED:
	for {
		for lo := input.Front(); lo != nil; lo = lo.Next() {
			mid := findIncreasingBlock(lo)
			if mid.Next() == nil && lo == input.Front() {
				break SORTED
			}
			hi := findIncreasingBlock(mid.Next())
			if hi == nil {
				continue
			}
			merge(input, lo, mid, hi)
			lo = hi
		}
	}
}

func findIncreasingBlock(lo *list.Element) (k *list.Element) {
	if lo == nil || lo.Next() == nil {
		return lo
	}

	for k = lo; k != nil && k.Next() != nil; k = k.Next() {
		if k.Value.(int) > k.Next().Value.(int) {
			break
		}
	}

	return k
}

func merge(input *list.List, lo, mid, hi *list.Element) {
	j := lo
	for k := mid.Next(); k != hi.Next() && j != mid; {
		nextK := k.Next()
		if k.Value.(int) < j.Value.(int) {
			input.MoveBefore(k, j)
			k = nextK
		} else {
			j = j.Next()
		}
	}
}

func printList(input *list.List) {
	for i := input.Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value.(int))
		if i.Next() != nil {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}

func toInt(input *list.List) (rsp []int) {
	for i := input.Front(); i != nil; i = i.Next() {
		if i.Next() != nil {
			rsp = append(rsp, i.Value.(int))
		}
	}
	return rsp
}
