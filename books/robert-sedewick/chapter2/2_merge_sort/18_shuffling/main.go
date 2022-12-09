package shuffling

import (
	"container/list"
	"math/rand"
)

// 2.2.18 Shuffling a linked list. Develop and implement a divide-and-conquer
// algorithm that randomly shuffles a linked list in linearithmic time and
// logarithmic extra space.

// NOTE: Both solutions need to transverse the array multiple times to find each
// randomly selected element
// NOTE: Review with peer, and review his approach

// Shuffle: Use another data structure to append the "randomly chosen" elements
// of the original input

func Shuffle(input *list.List) {
	if input == nil || input.Front() == nil || input.Front().Next() == nil {
		return
	}

	root := list.New()
	for input.Len() != 0 {
		// Generate a random number between [0,len(input)]
		idx := rand.Intn(input.Len())

		// Move to that position,
		n := input.Front()
		for j := 0; j < idx; j++ {
			n = n.Next()
		}

		// remove that element from the old list, put it on a Head new list
		root.PushBack(input.Remove(n))
	}

	// Swap the shuffled list by the now empty original list
	*input = *root
}

// ShuffleInPlace : Same as before, but without an extra data structure. Use it if
// you have a linked list with only a pointer to the start of the list
func ShuffleInPlace(input *list.List) {
	if input == nil || input.Front() == nil || input.Front().Next() == nil {
		return
	}

	processed := 0

	// Repeat until the numberProcessed == N
	for processed != input.Len() {
		// Find a random number on the remaining of the list [numberProcessed,N)
		idx := processed + rand.Intn(input.Len()-processed)

		// Move to that position
		n := input.Front()
		for i := 0; i < idx; i++ {
			n = n.Next()
		}

		// Move that element to the start
		input.MoveBefore(n, input.Front())
		processed++
	}
}

// ShuffleInPlaceInverse : Same as before, but without the strange index
// arithmetic; code is much clearer. Can be used if you have a pointer to both
// the start and end of the linked-list (usually won't be the case, unless double-linked list)
func ShuffleInPlaceInverse(input *list.List) {
	if input == nil || input.Front() == nil || input.Front().Next() == nil {
		return
	}

	notProcessed := input.Len()
	// Repeat until the numberProcessed == N
	for notProcessed != 0 {
		// Find a random number on the remaining of the list [numberProcessed,N)
		idx := rand.Intn(notProcessed)

		// Move to that position
		n := input.Front()
		for i := 0; i < idx; i++ {
			n = n.Next()
		}

		// Move that element to the end
		input.MoveAfter(n, input.Back())
		notProcessed--
	}
}
