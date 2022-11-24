package three_sum

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestThreeSum(t *testing.T) {
	rand.Seed(time.Now().UnixMilli())

	input := readFile(t, "./testdata/1Kints.txt")
	require.Equal(t, 70, ThreeSum(input))

	input = readFile(t, "./testdata/2Kints.txt")
	require.Equal(t, 528, ThreeSum(input))

	// n^3: 7.86s       n^2: 0.11s
	input = readFile(t, "./testdata/4Kints.txt")
	require.Equal(t, 4039, ThreeSum(input))

	// n^2: 0.43s
	input = readFile(t, "./testdata/8Kints.txt")
	require.Equal(t, 32074, ThreeSum(input))
}

func readFile(t *testing.T, filename string) []int {
	raw, err := os.ReadFile(filename)
	require.NoError(t, err)

	s := strings.ReplaceAll(string(raw), " ", "")
	lines := strings.Split(s, "\n")

	input := make([]int, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}

		n, err := strconv.Atoi(line)
		require.NoError(t, err)

		input = append(input, n)
	}
	return input
}
