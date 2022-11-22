package _9_remove_last_node

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRemoveLast(t *testing.T) {
	l := new(List[string])
	require.Empty(t, l.String())
	l.RemoveLast() // No Op
	require.Empty(t, l.String())
	l.Enqueue("hello")
	l.Enqueue("world")
	l.Enqueue("REMOVE ME")
	require.Equal(t, "hello -> world -> REMOVE ME", l.String())
	l.RemoveLast()
	require.Equal(t, "hello -> world", l.String())
	l.RemoveLast()
	l.RemoveLast()
	require.Empty(t, l.String())
}
