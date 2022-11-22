package main

import stack "exercises/books/robert-sedewick/chapter1/3_bags_queues_stacks/00_generic_stack"

func EvaluatePostfix(s string) int {

	result := new(stack.Stack[int])

	// 123*+4+
	for _, token := range s {
		switch token {
		case '+':
			b, a := result.Pop(), result.Pop()
			result.Push(a + b)
		case '-':
			b, a := result.Pop(), result.Pop()
			result.Push(a - b)
		case '*':
			b, a := result.Pop(), result.Pop()
			result.Push(a * b)
		case '/':
			b, a := result.Pop(), result.Pop()
			result.Push(a / b)
		default:
			result.Push(int(token - '0')) // Shitty integer-single-digit parser
		}

	}

	return result.Pop()

}
