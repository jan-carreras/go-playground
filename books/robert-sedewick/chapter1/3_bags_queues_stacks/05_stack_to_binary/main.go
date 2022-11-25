package main

import (
	"bytes"
	stack "github.com/jan-carreras/go-playground/books/robert-sedewick/chapter1/3_bags_queues_stacks/00_generic_stack"
)

// ToBinary converts 0 or positive integers to binary representation
func ToBinary(n int) string {
	if n < 0 {
		panic("negative values are not supported")
	}

	if n == 0 {
		return "0"
	}

	s := new(stack.Stack[int])
	for n > 0 {
		s.Push(n % 2)
		n /= 2
	}

	buf := bytes.Buffer{}
	for s.Length() > 0 {
		buf.WriteByte(byte('0' + s.Pop()))
	}

	return buf.String()
}
