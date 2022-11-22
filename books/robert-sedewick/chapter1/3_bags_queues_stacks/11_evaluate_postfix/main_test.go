package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEvaluatePostfix(t *testing.T) {
	require.Equal(t, 11, EvaluatePostfix("123*+4+"))
}
