package double_list_test

import (
	dl "github.com/jan-carreras/go-playground/books/robert-sedewick/chapter1/3_bags_queues_stacks/31_double_node"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestList_String(t *testing.T) {
	l := new(dl.List[string])
	checkList(t, l, "")

	l.InsertBeginning("hello")
	checkList(t, l, "hello")
}

func TestList_InsertBeginning(t *testing.T) {
	l := new(dl.List[string])
	l.InsertBeginning("world")
	l.InsertBeginning("hello")

	checkList(t, l, "hello -> world")
}

func TestList_RemoveBeginning(t *testing.T) {
	l := new(dl.List[string])
	l.InsertEnd("hello")
	l.InsertEnd("world")

	l.RemoveBeginning()
	checkList(t, l, "world")

	l.RemoveBeginning()
	checkList(t, l, "")
}

func TestList_InsertEnd(t *testing.T) {
	l := new(dl.List[string])
	l.InsertEnd("hello")
	l.InsertEnd("world")
	checkList(t, l, "hello -> world")
}

func TestList_RemoveEnd(t *testing.T) {
	l := new(dl.List[string])
	l.InsertEnd("hello")
	l.InsertEnd("world")

	l.RemoveEnd()
	checkList(t, l, "hello")

	l.RemoveEnd()
	checkList(t, l, "")
}

func TestList_InsertBefore(t *testing.T) {
	l := new(dl.List[string])
	l.InsertEnd("hello")
	l.InsertEnd("world")

	l.InsertBefore(dl.GetFirst(l), "START")
	checkList(t, l, "START -> hello -> world")

	l.InsertBefore(dl.GetLast(l), "END")
	checkList(t, l, "START -> hello -> END -> world")

	l.InsertBefore(dl.Next(dl.GetFirst(l)), "MID")
	checkList(t, l, "START -> MID -> hello -> END -> world")
}

func TestList_InsertAfter(t *testing.T) {
	l := new(dl.List[string])
	l.InsertEnd("hello")
	l.InsertEnd("world")

	l.InsertAfter(dl.GetFirst(l), "START")
	checkList(t, l, "hello -> START -> world")

	l.InsertAfter(dl.GetLast(l), "END")
	checkList(t, l, "hello -> START -> world -> END")

	l.InsertAfter(dl.Next(dl.GetFirst(l)), "MID")
	checkList(t, l, "hello -> START -> MID -> world -> END")
}

func TestList_Remove(t *testing.T) {
	l := new(dl.List[string])
	l.InsertEnd("hello")
	l.InsertEnd("dear")
	l.InsertEnd("world")

	l.Remove(dl.GetFirst(l))
	checkList(t, l, "dear -> world")

	l.Remove(dl.GetLast(l))
	checkList(t, l, "dear")

	l.InsertBeginning("hello")
	l.InsertEnd("world")

	checkList(t, l, "hello -> dear -> world")

	l.Remove(dl.Next(dl.GetFirst(l)))

	checkList(t, l, "hello -> world")
}

// checkList helper function that checks the list
func checkList(t *testing.T, l *dl.List[string], expected string) {
	expectedLength := 0
	if expected != "" {
		expectedLength = strings.Count(expected, " -> ") + 1
	}

	require.Equal(t, expected, l.String(), "traversal: forward")
	require.Equal(t, expected, dl.StringReverse(l), "traversal: backwards")
	require.Equal(t, expectedLength, l.Length())
}
