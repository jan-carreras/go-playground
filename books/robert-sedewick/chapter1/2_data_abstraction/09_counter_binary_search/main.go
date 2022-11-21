package main

// Instrument BinarySearch to use a Counter to count the total number of keys
// examined during all searches and then print the total after all searches are
// complete.
//
// Hint : Create a Counter in main() and pass it as an argument to rank().

type Counter int

func (c *Counter) Increment() {
	*c++
}

func (c *Counter) IncrementN(n int) {
	*c = Counter(n) + *c
}

func (c *Counter) Value() int {
	return int(*c)
}

func Rank(c *Counter, lst []int, search int) int {
	lo, hi := 0, len(lst)-1
	for lo <= hi {
		mid := lo + (hi - lo)
		switch {
		case search > lst[mid]:
			c.Increment()
			lo = mid + 1
		case search < lst[mid]:
			c.Increment()
			hi = mid - 1
		default:
			return mid
		}
	}

	return -1
}
