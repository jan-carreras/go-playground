package stack_test

import (
	stack "exercises/books/robert-sedewick/chapter1/3_bags_queues_stacks/00_generic_stack"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStack(t *testing.T) {

	s := stack.Stack[string]{}
	require.Equal(t, 0, s.Length())
	s.Push("world")
	require.Equal(t, 1, s.Length())
	s.Push("hello")
	require.Equal(t, 2, s.Length())

	require.Equal(t, "hello", s.Pop())
	require.Equal(t, 1, s.Length())
	require.Equal(t, "world", s.Pop())
	require.Equal(t, 0, s.Length())
}
