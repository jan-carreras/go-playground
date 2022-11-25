package main

import (
	adt "github.com/jan-carreras/go-playground/books/robert-sedewick/chapter1/3_bags_queues_stacks/adt"
)

func EvaluatePostfix(s string) int {
	result := adt.NewTypeStack[int]()

	// 123*+4+
	for _, token := range s {
		switch token {
		case '+':
			b, a := result.SPop(), result.SPop()
			result.Push(a + b)
		case '-':
			b, a := result.SPop(), result.SPop()
			result.Push(a - b)
		case '*':
			b, a := result.SPop(), result.SPop()
			result.Push(a * b)
		case '/':
			b, a := result.SPop(), result.SPop()
			result.Push(a / b)
		default:
			result.Push(int(token - '0')) // Shitty integer-single-digit parser
		}

	}

	return result.SPop()
}
