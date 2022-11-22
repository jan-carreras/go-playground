package main

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestCompute(t *testing.T) {
	input := strings.NewReader("( 1 + ( ( 2 + 3 ) * ( 4 * 5 ) ) )")
	result, err := Compute(input)
	require.NoError(t, err)
	require.Equal(t, 101.0, result)

	input = strings.NewReader("( ( 1 + ( ( 2 + 3 ) * ( 4 * 5 ) ) ) - 1 )")
	result, err = Compute(input)
	require.NoError(t, err)
	require.Equal(t, 100.0, result)
}
