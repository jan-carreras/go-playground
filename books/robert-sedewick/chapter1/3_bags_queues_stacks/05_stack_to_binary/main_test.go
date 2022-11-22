package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestToBinary(t *testing.T) {
	require.Equal(t, "110010", ToBinary(50))
	require.Equal(t, "10", ToBinary(2))
	require.Equal(t, "0", ToBinary(0))
}
