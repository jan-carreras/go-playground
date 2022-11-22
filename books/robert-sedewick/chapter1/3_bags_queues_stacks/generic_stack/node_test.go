package stack

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestNode(t *testing.T) {
	n := &Node[string]{}
	n.value = "hello"

	n2 := &Node[string]{}
	n2.value = "world"
	n.next = n2

	out := make([]string, 0)
	for ; n != nil; n = n.next {
		out = append(out, n.value)
	}

	require.Equal(t, "hello world", strings.Join(out, " "))
}
