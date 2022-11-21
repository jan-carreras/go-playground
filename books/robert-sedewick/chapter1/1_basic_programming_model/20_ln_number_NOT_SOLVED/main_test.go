package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestFactorial(t *testing.T) {

	require.EqualValues(t, 120.0, factorial(5))

	for i := 1.0; i < 10; i++ {
		println(int(i), int(factorial(i)), fmt.Sprintf("%.2f", math.Log(factorial(i))))
	}

}
