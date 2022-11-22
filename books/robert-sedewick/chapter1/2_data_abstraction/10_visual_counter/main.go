package main

// ✅ 1..2.10 Develop a class Counter that allows both increment and decrement
// operations.
//
// ✅ Take two arguments N and max in the constructor,
//
// where N specifies the maximum number of operations
//
// ✅and max specifies the maximum absolute value for the counter.
//
// ✅As a side effect, create a plot showing the value of the counter each time its
// tally changes.

type Counter struct {
	value  int
	values []int

	maxOperations int
	maxValue      int
}

func NewCounter(maxOperations, maxValue int) *Counter {
	return &Counter{
		values:        make([]int, 0),
		maxOperations: maxOperations,
		maxValue:      maxValue,
	}
}

func (c *Counter) Value() int {
	return c.value
}

func (c *Counter) Values() []int {
	return c.values
}

func (c *Counter) Increment() {
	c.add(1)
}

func (c *Counter) Decrement() {
	c.add(-1)
}

func (c *Counter) add(n int) {
	if c.maxOperations <= 0 {
		return // No Op. We have consumed all operations
	}

	c.maxOperations--
	if c.value+n > c.maxValue {
		return // We're already to the maximum. No Op.
	}

	c.value = c.value + n
	c.values = append(c.values, c.value)
}
