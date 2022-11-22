package generic_list_test

import (
	"exercises/books/robert-sedewick/chapter1/3_bags_queues_stacks/generic_list"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestList_List(t *testing.T) {
	l := generic_list.List[string]{}

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
