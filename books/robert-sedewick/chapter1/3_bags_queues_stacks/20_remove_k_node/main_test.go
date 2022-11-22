package _0_remove_k_node

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRemoveK(t *testing.T) {
	l := new(List[string])
	l.Enqueue("hello")
	l.Enqueue("my")
	l.Enqueue("stupid")
	l.Enqueue("world")
	require.Equal(t, "hello -> my -> stupid -> world", l.String())

	l.RemoveK(2)
	require.Equal(t, "hello -> my -> world", l.String())

	l.RemoveK(1)
	require.Equal(t, "hello -> world", l.String())

	l.RemoveK(1)
	require.Equal(t, "hello", l.String())

	l.Enqueue("world2")
	require.Equal(t, "hello -> world2", l.String())

	l.RemoveK(0)
	require.Equal(t, "world2", l.String())

	l.RemoveK(0)
	require.Equal(t, "", l.String())

	l.Enqueue("hello")
	l.Enqueue("world")
	require.Equal(t, "hello -> world", l.String())

	require.Equal(t, "hello", l.Dequeue())
	require.Equal(t, "world", l.Dequeue())

	require.Empty(t, l.String())
}
