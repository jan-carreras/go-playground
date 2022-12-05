package improvements

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestSort_AlreadySorted(t *testing.T) {
	input := []string{"A", "B", "C", "D", "E"}
	Sort(input)

	require.Equal(t, input, input)
}

func TestSort_NotSorted(t *testing.T) {
	rand.Seed(time.Now().UnixMilli())

	input := []string{}
	for i := 'a'; i <= 'z'; i++ {
		input = append(input, string(i))
	}
	expected := input

	for i := 0; i < len(input); i++ {
		a, b := rand.Intn(len(input)), rand.Intn(len(input))
		input[a], input[b] = input[b], input[a]
	}

	Sort(input)

	require.Equal(t, strings.Join(expected, ""), strings.Join(input, ""))
}
