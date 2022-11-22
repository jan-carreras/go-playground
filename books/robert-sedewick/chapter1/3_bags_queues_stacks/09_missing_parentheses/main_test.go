package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBalanceParentheses(t *testing.T) {
	expected := "( ( 1 + 2 ) * ( ( 3 - 4 ) * ( 5 - 6 ) ) )"
	having := BalanceParentheses("1 + 2 ) * 3 - 4 ) * 5 - 6 ) ) )")
	require.Equal(t, expected, having)
}
