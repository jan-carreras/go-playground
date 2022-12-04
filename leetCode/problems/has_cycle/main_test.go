package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHasCycle(t *testing.T) {
	zero := &ListNode{
		Val: 0,
	}

	two := &ListNode{
		Val:  2,
		Next: zero,
	}

	four := &ListNode{
		Val:  -4,
		Next: two,
	}
	zero.Next = four

	root := &ListNode{Val: 3, Next: two}
	require.True(t, hasCycle(root))
}
