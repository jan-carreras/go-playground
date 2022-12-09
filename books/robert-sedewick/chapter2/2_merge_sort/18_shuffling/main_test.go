package shuffling

import (
	"container/list"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestShuffle(t *testing.T) {
	rand.Seed(time.Now().UnixMilli())
	l := newList([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	Shuffle(l)
	fmt.Println(toInt(l))
}

func TestShuffleInPlace(t *testing.T) {
	rand.Seed(time.Now().UnixMilli())
	l := newList([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	ShuffleInPlace(l)
	fmt.Println(toInt(l))
}

func TestShuffleInPlaceReverse(t *testing.T) {
	rand.Seed(time.Now().UnixMilli())
	l := newList([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	ShuffleInPlaceInverse(l)
	fmt.Println(toInt(l))
}

func newList(input []int) *list.List {
	l := list.New()
	for _, value := range input {
		l.PushBack(value)
	}
	return l
}

func toInt(input *list.List) (rsp []int) {
	for i := input.Front(); i != nil; i = i.Next() {
		rsp = append(rsp, i.Value.(int))
	}
	return rsp
}
