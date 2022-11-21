package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBetweenRange(t *testing.T) {
	require.True(t, BetweenRange(0))
	require.True(t, BetweenRange(0.1))
	require.True(t, BetweenRange(0.9))
	require.True(t, BetweenRange(1))

	require.False(t, BetweenRange(-1))
	require.False(t, BetweenRange(1.00001))
}

func TestCheck(t *testing.T) {
	require.True(t, Check(0.5, 0.1))
	require.True(t, Check(0, 1))

	require.False(t, Check(0, 1.1))
	require.False(t, Check(-1, 0.5))
}

func TestCheckN(t *testing.T) {
	require.True(t, CheckN(0.5, 0.1))
	require.True(t, CheckN(0, 1))

	require.False(t, CheckN(0, 1.1))
	require.False(t, CheckN(-1, 0.5))
}
