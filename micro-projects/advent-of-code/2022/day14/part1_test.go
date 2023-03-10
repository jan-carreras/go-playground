package day14

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	f, err := os.Open("input.txt")
	require.NoError(t, err)

	err = part1(f)
	require.NoError(t, err)
}
