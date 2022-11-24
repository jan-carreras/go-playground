package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewQueue(t *testing.T) {

	q := &Queue{}
	q.Enqueue("hello")
	q.Enqueue("world")
	require.Equal(t, "hello -> world", q.String())

	r := NewQueue(q)
	require.Equal(t, "hello -> world", r.String())
	r.Enqueue("!!!")
	require.Equal(t, "hello -> world", q.String(), "old list is unchanged")
	require.Equal(t, "hello -> world -> !!!", r.String())

}
