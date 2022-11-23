package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFind(t *testing.T) {
	l := new(List[string])
	l.Enqueue("hello")
	l.Enqueue("world")

	require.True(t, l.Find("hello"))
	require.False(t, l.Find("cruel"))
	require.True(t, l.Find("world"))
}
