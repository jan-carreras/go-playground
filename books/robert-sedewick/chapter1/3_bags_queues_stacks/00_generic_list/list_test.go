package main_test

import (
	adt "github.com/jan-carreras/go-playground/books/robert-sedewick/chapter1/3_bags_queues_stacks/adt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestList_List(t *testing.T) {
	l := adt.TypedQueue[string]{}

	l.Enqueue("hello")
	l.Enqueue("world")

	require.Equal(t, "hello", l.SDequeue())
	require.Equal(t, "world", l.SDequeue())

	l.Enqueue("foo")
	l.Enqueue("bar")

	require.Equal(t, "foo", l.SDequeue())
	require.Equal(t, "bar", l.SDequeue())

	_, err := l.Dequeue()
	require.Error(t, err)
}
