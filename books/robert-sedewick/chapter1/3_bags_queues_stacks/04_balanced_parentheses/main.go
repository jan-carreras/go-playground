package main

import stack "exercises/books/robert-sedewick/chapter1/3_bags_queues_stacks/00_generic_stack"

func BalancedParentheses(input string) bool {
	openClose, closeOpen := openClosePairs()

	st := stack.Stack[rune]{}
	for _, r := range input {
		if _, ok := openClose[r]; ok {
			st.Push(r)
		} else if counter, ok := closeOpen[r]; ok {
			if st.Length() == 0 {
				return false
			}

			matching := st.Pop()
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
