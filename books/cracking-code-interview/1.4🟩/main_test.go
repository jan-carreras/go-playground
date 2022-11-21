package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPalindromePermutation(t *testing.T) {
	require.False(t, PalindromePermutation(""))
	require.True(t, PalindromePermutation("taco cat"))
	require.True(t, PalindromePermutation("atco cta"))
	require.True(t, PalindromePermutation("aa"))
	require.True(t, PalindromePermutation("aba"))     // aba is valid
	require.True(t, PalindromePermutation("baa"))     // aba is valid
	require.True(t, PalindromePermutation("abctcba")) // middle character is valid
	require.True(t, PalindromePermutation("abccba"))
	require.False(t, PalindromePermutation("abc"))
	require.False(t, PalindromePermutation("abca"))
}
