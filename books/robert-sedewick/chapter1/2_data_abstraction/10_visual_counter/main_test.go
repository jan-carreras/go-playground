package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVisualCounter(t *testing.T) {

	c := NewCounter(5, 2)
	require.Equal(t, 0, c.Value())
	c.Increment()
	require.Equal(t, 1, c.Value())
	c.Increment()
	require.Equal(t, 2, c.Value())
	c.Increment()
	require.Equal(t, 2, c.Value())
	c.Decrement()
	require.Equal(t, 1, c.Value())
	c.Decrement()
	require.Equal(t, 0, c.Value())
	c.Decrement()
	require.Equal(t, 0, c.Value()) // We don't have more operations left!

}
