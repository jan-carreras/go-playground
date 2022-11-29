package deck

import (
	"fmt"
	"math/rand"
	"time"
)

// 2.1.13 Deck sort. Explain how you would put a deck of cards in order by suit
// (in the order spades, hearts, clubs, diamonds) and by rank within each suit,
// with the restriction that the cards must be laid out face down in a row, and
// the only allowed operations are to check the values of two cards and to
// exchange two cards (keeping them face down).

type Deck []Card

func (d Deck) SelectSort() {
	N := len(d)
	for i := 0; i < N; i++ {
		min := i
		for j := i; j < N; j++ {
			if d.Less(j, min) {
				min = j
			}
		}
		d.Swap(i, min)
	}
}

func (d Deck) InsertSort() {
	N := len(d)
	for i := 1; i < N; i++ {
		for j := i; j > 0 && d.Less(j, j-1); j-- {
			d.Swap(j, j-1)
		}
	}
}

func (d Deck) Len() int {
	return len(d)
}

func (d Deck) Swap(x, y int) {
	d[x], d[y] = d[y], d[x]
}

func (d Deck) Less(x, y int) bool {
	a, b := d[x], d[y]
	if a.suit < b.suit {
		return true
	} else if a.suit == b.suit {
		if a.rank < b.rank {
			return true
		}
	}

	return false
}

func CreateDeck() Deck {
	cards := make(Deck, 0)
	for i := byte(0); i < 4; i++ {
		for j := byte(2); j <= 14; j++ {
			cards = append(cards, Card{suit: i, rank: j})
		}
	}
	return cards
}

func RandomizeDeck(deck Deck) {
	rand.Seed(time.Now().UnixMilli())
	for i := 0; i < len(deck); i++ {
		x, y := rand.Intn(len(deck)), rand.Intn(len(deck))
		deck[x], deck[y] = deck[y], deck[x]
	}
}

type Card struct {
	suit byte // spades, hearts, clubs, diamonds
	rank byte // 2,3,4,5,6,7,8,9,10,J,Q,K,A
}

func (c Card) String() string {
	return c.suitString() + c.rankString()

}

func (c Card) suitString() string {
	m := map[byte]string{0: "S", 1: "H", 2: "C", 3: "D"}
	if suit, ok := m[c.suit]; ok {
		return suit
	}

	panic("unknown suit")
}

func (c Card) rankString() string {
	if c.rank >= 2 && c.rank <= 10 {
		return fmt.Sprintf("%d", c.rank)
	}

	m := map[byte]string{11: "J", 12: "Q", 13: "K", 14: "A"}
	if rank, ok := m[c.rank]; ok {
		return rank
	}

	panic("unknown rank")
}
