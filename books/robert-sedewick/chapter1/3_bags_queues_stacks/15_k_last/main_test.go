package main

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestTail(t *testing.T) {
	input := `1`
	buf := &bytes.Buffer{}
	err := Tail(strings.NewReader(input), buf, 0)
	require.NoError(t, err)
	require.EqualValues(t, "the 0th line is: 1\n", buf.String())

	input = `1 boring line
2 awesome line`

	buf = &bytes.Buffer{}
	err = Tail(strings.NewReader(input), buf, 1)
	require.NoError(t, err)
	require.EqualValues(t, "the 1th line is: 1 boring line\n", buf.String())
}
