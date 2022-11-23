package main

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestEnqueue(t *testing.T) {
	q := new(CircularQueue[string])
	require.Equal(t, "", q.String())

	q.Enqueue("hello")
	require.Equal(t, "hello", q.String())

	q.Enqueue("world")
	require.Equal(t, "hello -> world", q.String())
}

func TestDequeue(t *testing.T) {
	q := new(CircularQueue[string])
	require.Equal(t, "", q.String())

	value, err := q.Dequeue()
	require.ErrorIs(t, ErrEmptyQueue, err)
	require.Empty(t, value)

	q.Enqueue("hello")
	q.Enqueue("world")
	value, err = q.Dequeue()
	require.NoError(t, err)
	require.Equal(t, "hello", value)
	value, err = q.Dequeue()
	require.NoError(t, err)
	require.Equal(t, "world", value)
	require.Zero(t, q.Length())
}

func TestCircularQueue_Each(t *testing.T) {
	q := new(CircularQueue[string])
	q.Enqueue("hello")
	q.Enqueue("world")

	values := make([]string, 0)
	q.Each(func(value string) {
		values = append(values, value)
	})

	require.Equal(t, "hello world", strings.Join(values, " "))
}
