package main

import (
	adt "github.com/jan-carreras/go-playground/books/robert-sedewick/chapter1/3_bags_queues_stacks/adt"
	"io"
	"math"
	"strconv"
	"strings"
)

// TODO: Re-do this exercise using generics!

// ( 1 + ( ( 2 + 3 ) * ( 4 * 5 ) ) )
func Compute(s io.Reader) (float64, error) {
	ops, oper := adt.NewTypeStack[string](), adt.NewTypeStack[float64]()

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
			op := ops.SPop()
			switch op {
			case "+":
				operB, operA := oper.SPop(), oper.SPop()
				oper.Push(operA + operB)
			case "-":
				operB, operA := oper.SPop(), oper.SPop()
				oper.Push(operA - operB)
			case "*":
				operB, operA := oper.SPop(), oper.SPop()
				oper.Push(operA * operB)
			case "/":
				divisor, dividend := oper.SPop(), oper.SPop()
				oper.Push(dividend / divisor)
			case "sqrt":
				o := oper.SPop()
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

	return oper.SPop(), nil
}
