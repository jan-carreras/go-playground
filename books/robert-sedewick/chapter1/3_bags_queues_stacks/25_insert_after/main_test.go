package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRemoveAfter(t *testing.T) {
	l := new(List[string])

	l.Enqueue("hello")
	l.Enqueue("world")

	require.Equal(t, "hello -> world", l.String())
	l.InsertAfter(l.first, &Node[string]{Value: "cruel"})
	require.Equal(t, "hello -> cruel -> world", l.String())
}
