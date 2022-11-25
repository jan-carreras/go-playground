package main

import (
	"github.com/jan-carreras/go-playground/books/robert-sedewick/chapter1/3_bags_queues_stacks/adt"
)

// 1.3.21 Write a method find() that takes a linked list and a string key as
// arguments and returns true if some node in the list has key as its item field,
// false otherwise.

type List[T comparable] struct {
	adt.TypedQueue[T]
}

func (l *List[T]) Find(search T) bool {
	found := false
	l.EachN(func(_ int, v any) bool {
		found = v == search
		return !found // Iterate until we find search
	})

	return found
}
