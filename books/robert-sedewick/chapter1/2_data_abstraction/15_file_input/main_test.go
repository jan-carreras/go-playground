package main

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestReadInts(t *testing.T) {
	input := strings.NewReader("123 4 5 6 7777")

	ints, err := ReadInts(input)
	require.NoError(t, err)
	require.Equal(t, []int{123, 4, 5, 6, 7777}, ints)

	input = strings.NewReader("")
	ints, err = ReadInts(input)
	require.NoError(t, err)
	require.Empty(t, ints)

	input = strings.NewReader("1 2 3 aaa")
	_, err = ReadInts(input)
	require.Error(t, err)

}
