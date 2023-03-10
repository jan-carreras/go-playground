package day15

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Test(t *testing.T) {
	f, err := os.Open("input.txt")
	require.NoError(t, err)
	t.Cleanup(func() { f.Close() })

	err = part1(f)
	require.NoError(t, err)

}
