package brackets

var opening = map[rune]bool{'[': true, '{': true, '(': true}
var pairs = map[rune]rune{']': '[', ')': '(', '}': '{'}

func Bracket(s string) bool {
	stack := make([]rune, 0)

	for _, r := range s {
		if _, isOpeningRune := opening[r]; isOpeningRune {
			stack = append(stack, r)
			continue
		}
		if expectedPair, isClosingRune := pairs[r]; isClosingRune {
			if len(stack) == 0 && stack[len(stack)-1] != expectedPair {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
