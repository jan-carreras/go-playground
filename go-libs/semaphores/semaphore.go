package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"log"
	"sync"
)

func main2() {

	workToDo := 100
	maxConcurrency := 10
	activeGoroutines := 0

	tokens := make(chan interface{}, maxConcurrency)

	wg := sync.WaitGroup{}

	for i := 0; i < workToDo; i++ {
		wg.Add(1)
		tokens <- true
		activeGoroutines++
		go func(i int) {
			fmt.Printf("[%d] Number of active goroutines: %d\n", i, activeGoroutines)
			<-tokens
			activeGoroutines--
			wg.Done()
		}(i)
	}

}

func main() {
	ctx := context.Background()

	const workToDo = 100
	const maxConcurrency = 10
	activeGoroutines := 0

	sem := semaphore.NewWeighted(maxConcurrency)

	wg := sync.WaitGroup{}

	for i := 0; i < workToDo; i++ {
		wg.Add(1)
		if err := sem.Acquire(ctx, 1); err != nil {
			log.Fatalln(err)
		}

		activeGoroutines++
		go func(i int) {
			fmt.Printf("[%d] Number of active goroutines: %d\n", i, activeGoroutines)
			activeGoroutines--
			sem.Release(1)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
