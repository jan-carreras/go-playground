package deck_test

import (
	"fmt"
	deck "github.com/jan-carreras/go-playground/books/robert-sedewick/chapter2/1_elementary_sorts/13_deck"
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

func TestDeck(t *testing.T) {
	d := deck.CreateDeck()
	ordered := fmt.Sprintf("%v", d)

	t.Run("deck.SelectSort", func(t *testing.T) {
		deck.RandomizeDeck(d)
		d.SelectSort()
		require.Equal(t, ordered, fmt.Sprintf("%v", d))
	})
	t.Run("deck.InsertSort", func(t *testing.T) {
		deck.RandomizeDeck(d)
		d.InsertSort()
		require.Equal(t, ordered, fmt.Sprintf("%v", d))
	})

	t.Run("sort.Slice", func(t *testing.T) {
		deck.RandomizeDeck(d)
		sort.Slice(d, d.Less)
		require.Equal(t, ordered, fmt.Sprintf("%v", d))
	})

	t.Run("sort.Sort", func(t *testing.T) {
		deck.RandomizeDeck(d)
		sort.Sort(d)
		require.Equal(t, ordered, fmt.Sprintf("%v", d))
	})
}
