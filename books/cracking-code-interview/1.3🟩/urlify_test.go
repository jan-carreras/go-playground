package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUrlify(t *testing.T) {
	require.Equal(t, []byte(""), Urlify([]byte(""), 0))
	require.Equal(t, []byte("a"), Urlify([]byte("a"), 1))
	require.Equal(t, []byte("abc"), Urlify([]byte("abc"), 3))
	require.Equal(t, []byte("abc%20"), Urlify([]byte("abc   "), 4))
	require.Equal(t, []byte("%20abc"), Urlify([]byte(" abc  "), 4))
	require.Equal(t, []byte("Mr%20John%20Smith"), Urlify([]byte("Mr John Smith    "), 13))
}
