package main

import (
	"fmt"
	"sync"
	"time"
)

// From: https://blog.janestreet.com/what-a-jane-street-dev-interview-is-like/
//
// Have you heard about memoization? Can you carefully describe what it is? If
// you haven’t heard about it, don’t worry. We’ll bring you up to speed. (A good
// introduction is on Wikipedia.)
//
// Now let’s say there is a function f of type int -> int whose output only
// depends on the input. f is very expensive to compute. We’d like you to write a
// memoized version of this function, i.e. another function g of the same type,
// that returns the same values – g(x) = f(x) for all x – but only does the
// expensive computation once for each input value.

func main() {
	mf := memoizedF()

	for i := 0; i < 10; i++ {
		n := mf(1)
		fmt.Println(n)

		n = mf(2)
		fmt.Println(n)
	}

}

func memoizedF() func(n int) int {
	mem := make(map[int]int)
	mux := sync.RWMutex{}
	return func(n int) int {
		mux.RLock()
		value, ok := mem[n]
		mux.RUnlock()

		if ok {
			return value
		}

		// If not memoized before, compute the value
		value = f(n)

		mux.Lock()
		mem[n] = value
		mux.Unlock()

		return value
	}
}

func f(n int) int {
	time.Sleep(time.Second)
	return n * 10
}
