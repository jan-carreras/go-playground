package faster_merge

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestMerge(t *testing.T) {
	input := []string{"D", "C", "A", "B"}
	merge[string](input, 0, len(input)/2, len(input)-1)

	require.Equal(t, "ABDC", strings.Join(input, ""))

	fmt.Println(input)
}

func TestMerge2(t *testing.T) {
	input := []string{"3", "4", "1", "2"}
	merge2[string](input, 0, len(input)/2, len(input)-1)

	require.Equal(t, "1234", strings.Join(input, ""))

	fmt.Println(input)
}
