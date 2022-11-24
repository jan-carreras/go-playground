package _0_remove_k_node

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRemoveN(t *testing.T) {
	l := new(List[string])
	l.Enqueue("hello")
	l.Enqueue("my")
	l.Enqueue("stupid")
	l.Enqueue("world")
	require.Equal(t, "hello -> my -> stupid -> world", l.String())

	l.RemoveN(2)
	require.Equal(t, "hello -> my -> world", l.String())

	l.RemoveN(1)
	require.Equal(t, "hello -> world", l.String())

	l.RemoveN(1)
	require.Equal(t, "hello", l.String())

	l.Enqueue("world2")
	require.Equal(t, "hello -> world2", l.String())

	l.RemoveN(0)
	require.Equal(t, "world2", l.String())

	l.RemoveN(0)
	require.Equal(t, "", l.String())

	l.Enqueue("hello")
	l.Enqueue("world")
	require.Equal(t, "hello -> world", l.String())

	require.Equal(t, "hello", l.SDequeue())
	require.Equal(t, "world", l.SDequeue())

	require.Empty(t, l.String())
}
