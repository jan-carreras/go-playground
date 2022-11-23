package steque_test

import (
	steque "exercises/books/robert-sedewick/chapter1/3_bags_queues_stacks/32_steque"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSteque(t *testing.T) {
	s := new(steque.Steque[string])

	s.Push("world")
	s.Push("hello")

	require.Equal(t, "hello -> world", s.String())

	s.Enqueue("END")
	require.Equal(t, "hello -> world -> END", s.String())

	v, err := s.Pop()
	require.NoError(t, err)
	require.Equal(t, "hello", v)
	require.Equal(t, "world -> END", s.String())

	_, _ = s.Pop()
	_, _ = s.Pop()
	require.Equal(t, "", s.String())

	v, err = s.Pop()
	require.ErrorIs(t, err, steque.ErrEmpty)
	require.Empty(t, v)
}
