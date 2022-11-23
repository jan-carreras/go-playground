package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRemoveAll(t *testing.T) {
	l := new(List[string])

	l.Enqueue("hello")
	l.Enqueue("hello")
	l.RemoveAll("hello")
	require.Equal(t, "", l.String())

	l.Enqueue("hello")
	l.Enqueue("hello")
	l.Enqueue("world")
	l.RemoveAll("hello")
	require.Equal(t, "world", l.String())

	l.Enqueue("hello")
	l.Enqueue("world")
	require.Equal(t, "world -> hello -> world", l.String())
	l.RemoveAll("hello")
	require.Equal(t, "world -> world", l.String())

	l.Enqueue("hello")
	l.Enqueue("hello")
	l.RemoveAll("hello")
	require.Equal(t, "world -> world", l.String())
	l.Enqueue("world")
	require.Equal(t, "world -> world -> world", l.String())

}
