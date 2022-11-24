package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCopyStack(t *testing.T) {
	stack := &S{}
	stack.Push("world")
	stack.Push("hello")
	require.Equal(t, "hello -> world", stack.String())

	newStack := CopyStack(stack)
	newStack.Push(">>>")
	require.Equal(t, "hello -> world", stack.String(), "old stack is the same")
	require.Equal(t, ">>> -> hello -> world", newStack.String(), "old stack is the same")
}
