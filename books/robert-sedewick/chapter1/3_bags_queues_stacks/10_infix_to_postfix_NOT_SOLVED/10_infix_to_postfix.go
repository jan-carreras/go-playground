package main

import (
	stack "exercises/books/robert-sedewick/chapter1/3_bags_queues_stacks/00_generic_stack"
	"fmt"
	"strings"
)

// 1. Priority of the operations for grouping:
//		0: * /     21: + -
//
//
// A + B * C + D   -->   ABC*+D+

func InfixToPostfix(s string) string {
	operator, operand := new(stack.Stack[string]), new(stack.Stack[string])

	for _, token := range strings.Split(s, " ") {
		switch token {
		case "*", "+", "-", "/":
			operator.Push(token)
		default:
			operand.Push(token)
			/*if operator.Length() == 0 {
				continue
			}*/

			/*op := operator.Pop()*/
			/*if priority == 0 && (op == "*" || op == "/") {
				b, a := operand.Pop(), operand.Pop()
				fmt.Println(priority, fmt.Sprintf("%s%s%s", a, b, op))

				operand.Push(fmt.Sprintf("%s%s%s", a, b, op))
			} else if priority == 1 && (op == "+" || op == "-") {
				b, a := operand.Pop(), operand.Pop()
				fmt.Println(priority, fmt.Sprintf("%s%s%s", a, b, op))
				operand.Push(fmt.Sprintf("%s%s%s", a, b, op))
			} else {
				operator.Push(op)
			}*/

			// Check priority, to see if we need to mash something together
		}
	}

	fmt.Println(operator)
	fmt.Println(operand)

	return operand.Pop()
}
