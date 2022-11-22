package main_test

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestList_List(t *testing.T) {
	l := _0_generic_list.List[string]{}

	l.Enqueue("hello")
	l.Enqueue("world")

	require.Equal(t, "hello", l.Dequeue())
	require.Equal(t, "world", l.Dequeue())

	l.Enqueue("foo")
	l.Enqueue("bar")

	require.Equal(t, "foo", l.Dequeue())
	require.Equal(t, "bar", l.Dequeue())

	require.PanicsWithValue(t, "queue is empty", func() {
		l.Dequeue()
	})
}
