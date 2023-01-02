package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestCountContained(t *testing.T) {
	f, err := os.Open("./input.txt")
	require.NoError(t, err)

	total, err := countIsContained(f)
	require.NoError(t, err)

	require.Equal(t, 584, total)
}

func TestCountOverlaps(t *testing.T) {
	f, err := os.Open("./input.txt")
	require.NoError(t, err)

	total, err := countOverlaps(f)
	require.NoError(t, err)

	require.Equal(t, 933, total)
}

func TestContained(t *testing.T) {

	t.Run("does not overlap", func(t *testing.T) {
		a := assignement{start: 1, end: 2}
		b := assignement{start: 3, end: 4}
		require.False(t, a.isContained(b))
		require.False(t, b.isContained(a))
	})

	t.Run("b is contained in a", func(t *testing.T) {
		a := assignement{start: 1, end: 10}
		b := assignement{start: 3, end: 4}
		require.False(t, a.isContained(b))
		require.True(t, b.isContained(a))
	})

	t.Run("equal ranges", func(t *testing.T) {
		a := assignement{start: 1, end: 10}
		require.True(t, a.isContained(a))
	})
}

func TestOverlaps(t *testing.T) {
	t.Run("does not overlap", func(t *testing.T) {
		a := assignement{start: 1, end: 2}
		b := assignement{start: 3, end: 4}
		require.False(t, a.overlaps(b))
		require.False(t, b.overlaps(a))
	})

	t.Run("b overlaps a", func(t *testing.T) {
		a := assignement{start: 1, end: 3}
		b := assignement{start: 3, end: 4}
		require.True(t, a.overlaps(b))
		require.True(t, b.overlaps(a))
	})

	t.Run("b is inside a", func(t *testing.T) {
		a := assignement{start: 1, end: 10}
		b := assignement{start: 3, end: 4}
		require.True(t, a.overlaps(b))
		require.True(t, b.overlaps(a))
	})
}
