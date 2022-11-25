package main

import (
	"github.com/jan-carreras/go-playground/books/robert-sedewick/chapter1/3_bags_queues_stacks/adt"
)

func BalancedParentheses(input string) bool {
	openClose, closeOpen := openClosePairs()

	st := adt.NewTypeStack[rune]()

	for _, r := range input {
		if _, ok := openClose[r]; ok {
			st.Push(r)
		} else if counter, ok := closeOpen[r]; ok {
			if st.Len() == 0 {
				return false
			}

			matching := st.SPop()
			if counter != matching {
				return false
			}
		} else {
			return false // Unexpected character found
		}
	}

	return true
}

type pair map[rune]rune

func openClosePairs() (open pair, close pair) {
	// Pairs in: open/close
	openClose := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}

	// Reversed data structure
	closeOpen := make(map[rune]rune)
	for k, v := range openClose {
		closeOpen[v] = k
	}

	return openClose, closeOpen
}
