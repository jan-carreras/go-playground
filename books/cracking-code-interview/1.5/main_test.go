package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDistance_ZeroEdits(t *testing.T) {
	require.True(t, Distance("pale", "pale"))
}

func TestDistance_TooDistant(t *testing.T) {
	require.False(t, Distance("pale", "p"))
	require.False(t, Distance("pale", "bake"))
	require.False(t, Distance("pale", "bakebake"))
}

func TestDistance_Replace(t *testing.T) {
	require.True(t, Distance("pale", "pala"))
	require.False(t, Distance("pale", "paaa"))
}

func TestDistance_Examples(t *testing.T) {
	require.True(t, Distance("pale", "ple"))
	require.True(t, Distance("pales", "pale"))
	require.True(t, Distance("pale", "pales"))
	require.True(t, Distance("pale", "bale"))
	require.False(t, Distance("pale", "bake"))
}
