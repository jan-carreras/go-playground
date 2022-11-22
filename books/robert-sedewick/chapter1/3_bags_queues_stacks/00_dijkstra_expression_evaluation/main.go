package main

import (
	"github.com/golang-collections/collections/stack"
	"io"
	"math"
	"strconv"
	"strings"
)

// TODO: Re-do this exercise using generics!

// ( 1 + ( ( 2 + 3 ) * ( 4 * 5 ) ) )
func Compute(s io.Reader) (float64, error) {
	ops, oper := stack.New(), stack.New()
	input, err := io.ReadAll(s)
	if err != nil {
		return 0, err
	}

	for _, c := range strings.Split(string(input), " ") {
		switch c {
		case "(": // Ignore it
		case "+", "-", "*", "/", "sqrt":
			ops.Push(c)
		case ")": // We can perform the operation!
			op := ops.Pop().(string)
			switch op {
			case "+":
				operB, operA := oper.Pop().(float64), oper.Pop().(float64)
				oper.Push(operA + operB)
			case "-":
				operB, operA := oper.Pop().(float64), oper.Pop().(float64)
				oper.Push(operA - operB)
			case "*":
				operB, operA := oper.Pop().(float64), oper.Pop().(float64)
				oper.Push(operA * operB)
			case "/":
				divisor, dividend := oper.Pop().(float64), oper.Pop().(float64)
				oper.Push(dividend / divisor)
			case "sqrt":
				o := oper.Pop().(float64)
				oper.Push(math.Sqrt(o))
			}

		default:
			number, err := strconv.ParseFloat(c, 64)
			if err != nil {
				return 0, err
			}

			oper.Push(number)
		}
	}

	return oper.Pop().(float64), nil
}
