package day16

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Test(t *testing.T) {
	f, err := os.Open("input.txt")
	require.NoError(t, err)
	err = day1(f)
	require.NoError(t, err)
}
