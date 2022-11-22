package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBalancedParentheses(t *testing.T) {
	require.False(t, BalancedParentheses("[(a)]"))
	require.True(t, BalancedParentheses("[(<>)]{}{[()()]()}"))
	require.False(t, BalancedParentheses("[(])"))
}
