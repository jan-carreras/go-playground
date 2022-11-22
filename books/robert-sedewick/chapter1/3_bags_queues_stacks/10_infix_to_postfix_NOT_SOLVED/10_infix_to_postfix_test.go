package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInfixToPostfix(t *testing.T) {
	require.Equal(t, "ABC*+D+", InfixToPostfix("A + B * C + D"))
	//require.Equal(t, "B+CDE/*-F+", InfixToPostfix("( (A + B) – C * (D / E) ) + F"))
}
