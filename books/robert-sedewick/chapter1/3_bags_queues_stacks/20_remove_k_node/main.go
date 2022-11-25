package _0_remove_k_node

import (
	adt "github.com/jan-carreras/go-playground/books/robert-sedewick/chapter1/3_bags_queues_stacks/adt"
)

// 1.3.20 Write a method delete() that takes an int argument k and deletes the
// kth element in a linked list, if it exists.

type List[T any] struct {
	adt.TypedQueue[T]
}
