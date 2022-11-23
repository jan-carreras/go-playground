package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMax(t *testing.T) {
	l := new(List[int])
	l.Enqueue(1)
	l.Enqueue(2)
	l.Enqueue(3)
	l.Enqueue(2)
	l.Enqueue(5)
	l.Enqueue(1)

	require.Equal(t, 5, Max(l))
}

func TestMax_Empty(t *testing.T) {
	l := new(List[int])
	require.Equal(t, 0, Max(l))
}
