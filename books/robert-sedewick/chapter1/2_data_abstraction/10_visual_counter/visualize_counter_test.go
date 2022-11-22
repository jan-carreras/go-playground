package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVisualizeCounter(t *testing.T) {
	c := NewCounter(10, 10)
	c.Increment()
	c.Increment()
	c.Increment()
	c.Decrement()
	c.Decrement()
	c.Increment()
	c.Increment()
	c.Increment()
	c.Increment()
	c.Increment()
	fmt.Println(c.Values())
	require.NoError(t, VisualizeCounter(c))
}
