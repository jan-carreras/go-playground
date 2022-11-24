package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMoveToFront(t *testing.T) {
	m := MoveToFront{}
	m.Add("hello")
	m.Add("world")
	require.Equal(t, "world -> hello -> ", m.String())
	m.Add("hello")
	require.Equal(t, "hello -> world -> ", m.String())
}
