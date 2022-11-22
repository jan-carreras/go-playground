package main

import (
	stack "exercises/books/robert-sedewick/chapter1/3_bags_queues_stacks/00_generic_stack"
	"fmt"
	"strings"
)

//  1.3.9 Write a program that takes from standard input an expression without
//  left parentheses and prints the equivalent infix expression with the
//  parentheses inserted.
//
// For example, given the input:
//
//		1 + 2 ) * 3 - 4 ) * 5 - 6 ) ) )
//
// your program should print:
//
//		( ( 1 + 2 ) * ( ( 3 - 4 ) * ( 5 - 6 ) )
//

func BalanceParentheses(input string) string {
	operators := new(stack.Stack[string])
	operands := new(stack.Stack[string])

	for _, token := range strings.Split(input, " ") {
		switch token {
		case "*", "-", "+", "/":
			operators.Push(token)
		case ")":
			op := operators.Pop()
			b, a := operands.Pop(), operands.Pop()
			operands.Push(fmt.Sprintf("( %s %s %s )", a, op, b))
		// Balance parenthesis
		default:
			operands.Push(token)
		}
	}

	return operands.Pop()
}
