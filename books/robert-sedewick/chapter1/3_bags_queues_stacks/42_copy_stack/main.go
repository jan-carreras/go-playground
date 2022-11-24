package main

import (
	"exercises/books/robert-sedewick/chapter1/3_bags_queues_stacks/00_adt"
)

// 1.3.42 Copy a stack. Create a new constructor for the linked-list implementation of Stack so that
//
// 		Stack<Item> t = new Stack<Item>(s);
//
// makes t a reference to a new and independent copy of the stack s.

type S struct {
	adt.Stack
}

func CopyStack(oldStack *S) (newStack *S) {
	reversed := &S{}
	for oldStack.Len() != 0 {
		v, err := oldStack.Pop()
		if err != nil {
			panic("stack modified when performing operation")
		}

		reversed.Push(v)
	}

	newStack = &S{}
	for reversed.Len() != 0 {
		v, _ := reversed.Pop()
		oldStack.Push(v)
		newStack.Push(v)
	}

	return newStack
}
