package editor

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBuffer_Insert(t *testing.T) {
	b := &Buffer{}

	b.Insert('h')
	b.Insert('i')
	require.Equal(t, "hi|", b.String())
}

func TestBuffer_Left(t *testing.T) {
	b := &Buffer{}

	b.Insert('h')
	b.Insert('i')
	require.Equal(t, "hi|", b.String())

	b.Left(1)
	require.Equal(t, "h|i", b.String())

	b.Left(1)
	require.Equal(t, "|hi", b.String())

	b.Left(100)
	require.Equal(t, "|hi", b.String())
}

func TestBuffer_Right(t *testing.T) {
	b := &Buffer{}

	b.Insert('h')
	b.Insert('i')
	b.Left(100)
	require.Equal(t, "|hi", b.String())

	b.Right(1)
	require.Equal(t, "h|i", b.String())

	b.Right(1)
	require.Equal(t, "hi|", b.String())

	b.Right(100)
	require.Equal(t, "hi|", b.String())
}

func TestBuffer_Remove(t *testing.T) {
	b := &Buffer{}
	b.Insert('h')
	b.Insert('i')
	require.Equal(t, "hi|", b.String())

	c, err := b.Delete()
	require.NoError(t, err)
	require.Equal(t, 'i', c)
	require.Equal(t, "h|", b.String())

	c, err = b.Delete()
	require.NoError(t, err)
	require.Equal(t, 'h', c)

	c, err = b.Delete()
	require.ErrorIs(t, ErrNothingToDelete, err)
	require.Equal(t, rune(0), c)
}
