package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBinary(t *testing.T) {
	require.Equal(t, "10", Binary(2))
	require.Equal(t, "1000", Binary(8))
	require.Equal(t, "110", Binary(6))
}
