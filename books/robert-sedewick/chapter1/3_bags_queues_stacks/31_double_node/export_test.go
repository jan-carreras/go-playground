package double_list

import (
	"fmt"
	"strings"
)

// StringReverse returns the same output as String (start -> end), but iterating backwards
// used to make sure that the Node.previous links are properly linked
// Warning: Used only for testing purposed! This function is not in the final binary!
func StringReverse(l *List[string]) string {
	if l.length == 0 {
		return ""
	}

	lst := make([]string, l.length)
	eachNReverse(l, func(n int, value string) {
		if n >= l.length {
			panic(fmt.Sprintf("unexpected index. length=%d index=%d . index should be always lower than length", l.length, n))
		}
		lst[n] = value
	})

	b := strings.Builder{}
	for n, value := range lst {
		b.WriteString(fmt.Sprintf("%v", value))
		if n != l.length-1 { // Last element on the list
			b.WriteString(" -> ")
		}
	}

	return b.String()
}

// eachNReverse calls fnx for each element, passing its position (zero-indexed) and the value
// but from end to start
// Warning: Used only for testing purposed! This function is not in the final binary!
func eachNReverse(l *List[string], fnx func(n int, value string)) {
	n := l.last
	for i := l.Length() - 1; n != nil; i-- {
		fnx(i, n.value)
		n = n.previous
	}
}

func GetFirst(l *List[string]) *Node[string] {
	return l.first
}

func GetLast(l *List[string]) *Node[string] {
	return l.last
}

func Next(n *Node[string]) *Node[string] {
	return n.next
}
