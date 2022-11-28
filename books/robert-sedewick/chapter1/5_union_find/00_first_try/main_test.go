package main

import (
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	input := strings.NewReader(readFile(t, "./testdata/input.txt"))

	err := run(input)
	require.NoError(t, err)
}

func TestUnion(t *testing.T) {
	uf := NewUF(10)
	require.Equal(t, 10, uf.Count())
	uf.Union(4, 3)
	require.Equal(t, 9, uf.Count())
	uf.Union(3, 8)
	require.Equal(t, 8, uf.Count())
	uf.Union(8, 3)
	require.Equal(t, 8, uf.Count())
	uf.Union(9, 4)
	require.Equal(t, 7, uf.Count())

}

func readFile(t *testing.T, filename string) string {
	f, err := os.Open(filename)
	require.NoError(t, err)

	data, err := io.ReadAll(f)
	require.NoError(t, err)

	return string(data)
}
