package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Test(t *testing.T) {
	f, err := os.Open("input.txt")
	require.NoError(t, err)

	result, err := run(f)
	require.NoError(t, err)

	require.Equal(t, 16060, result)

}
