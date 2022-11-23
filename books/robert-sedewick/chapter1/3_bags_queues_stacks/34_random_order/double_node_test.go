package random_order

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestList_String(t *testing.T) {
	l := new(List[string])
	checkList(t, l, "")

	l.insertBeginning("hello")
	checkList(t, l, "hello")
}

func TestList_InsertBeginning(t *testing.T) {
	l := new(List[string])
	l.insertBeginning("world")
	l.insertBeginning("hello")

	checkList(t, l, "hello -> world")
}

func TestList_RemoveBeginning(t *testing.T) {
	l := new(List[string])
	l.insertEnd("hello")
	l.insertEnd("world")

	l.removeBeginning()
	checkList(t, l, "world")

	l.removeBeginning()
	checkList(t, l, "")
}

func TestList_InsertEnd(t *testing.T) {
	l := new(List[string])
	l.insertEnd("hello")
	l.insertEnd("world")
	checkList(t, l, "hello -> world")
}

func TestList_RemoveEnd(t *testing.T) {
	l := new(List[string])
	l.insertEnd("hello")
	l.insertEnd("world")

	l.removeEnd()
	checkList(t, l, "hello")

	l.removeEnd()
	checkList(t, l, "")
}

func TestList_InsertBefore(t *testing.T) {
	l := new(List[string])
	l.insertEnd("hello")
	l.insertEnd("world")

	l.insertBefore(getFirst(l), "START")
	checkList(t, l, "START -> hello -> world")

	l.insertBefore(getLast(l), "END")
	checkList(t, l, "START -> hello -> END -> world")

	l.insertBefore(Next(getFirst(l)), "MID")
	checkList(t, l, "START -> MID -> hello -> END -> world")
}

func TestList_InsertAfter(t *testing.T) {
	l := new(List[string])
	l.insertEnd("hello")
	l.insertEnd("world")

	l.insertAfter(getFirst(l), "START")
	checkList(t, l, "hello -> START -> world")

	l.insertAfter(getLast(l), "END")
	checkList(t, l, "hello -> START -> world -> END")

	l.insertAfter(Next(getFirst(l)), "MID")
	checkList(t, l, "hello -> START -> MID -> world -> END")
}

func TestList_Remove(t *testing.T) {
	l := new(List[string])
	l.insertEnd("hello")
	l.insertEnd("dear")
	l.insertEnd("world")

	l.remove(getFirst(l))
	checkList(t, l, "dear -> world")

	l.remove(getLast(l))
	checkList(t, l, "dear")

	l.insertBeginning("hello")
	l.insertEnd("world")

	checkList(t, l, "hello -> dear -> world")

	l.remove(Next(getFirst(l)))

	checkList(t, l, "hello -> world")
}

// checkList helper function that checks the list
func checkList(t *testing.T, l *List[string], expected string) {
	expectedLength := 0
	if expected != "" {
		expectedLength = strings.Count(expected, " -> ") + 1
	}

	require.Equal(t, expected, l.String(), "traversal: forward")
	require.Equal(t, expected, stringReverse(l), "traversal: backwards")
	require.Equal(t, expectedLength, l.length)
}

// stringReverse returns the same output as String (start -> end), but iterating backwards
// used to make sure that the Node.previous links are properly linked
// Warning: Used only for testing purposed! This function is not in the final binary!
func stringReverse(l *List[string]) string {
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
	for i := l.length - 1; n != nil; i-- {
		fnx(i, n.value)
		n = n.previous
	}
}

func getFirst(l *List[string]) *node[string] {
	return l.first
}

func getLast(l *List[string]) *node[string] {
	return l.last
}

func Next(n *node[string]) *node[string] {
	return n.next
}
