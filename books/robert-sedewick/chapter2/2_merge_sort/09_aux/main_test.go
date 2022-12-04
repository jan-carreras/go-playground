package aux

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"sort"
	"strings"
	"testing"
)

func TestSort(t *testing.T) {
	output := new(bytes.Buffer)
	s := new(TopDown[string]).WithDebug(output)

	input := strings.Split("E A S Y Q U E S T I O N", " ")

	s.Sort(input)

	inputCopy := input
	sort.Strings(inputCopy)
	require.Equal(t, inputCopy, input)
}
