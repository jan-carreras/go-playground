// Producer-consumer exercise
// Please complete the "consume" function below

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Workload is a general interface for a computational task
type Workload interface {
	Process() error
}

// Producer returns a new workload for as long as 'done' returns false.
// If the producer returns `true`, it must not be used anymore.
type Producer func() (workload Workload, done bool)

// consume calls `producer` on all workloads.
// `workers` configures the number of active go-routines.
// The producer is used until it signals it has no more values available.
// The number of items produced is unknown to consume.
//
// example: consume(1, testProducer(5))
//
//	Processing: 0
//	Processing: 1
//	Processing: 2
//	Processing: 3
//	Processing: 4
//
// Requirements:
// - workers parameter sets the number of go-routines that must be used.
// - consume returns only after all workloads have been processed
// - calls to producer are not thread-safe
// - The go-routines must be stopped before consume returns
// - workers should print log statements on start/stop
func consume(workers int, producer Producer) {
	// Input validation
	if workers <= 0 {
		return
	}

	// Define a waiting mechanism
	wg := &sync.WaitGroup{}
	wg.Add(workers)

	// Define a mechanism to pass work to do into the goroutines
	inbox := make(chan Workload, workers)
	errors := make(chan error, workers)

	drainChannel := false
	// Start the goroutines
	for i := 0; i < workers; i++ {
		go worker(i, wg, inbox, errors)
	}

	// Read from the producer until we don't have anything else to do
MAINLOOP:
	for {
		workToDo, done := producer()
		if done {
			break MAINLOOP
		}

		// This might block if the channel is full — it's fine for as long
		// as there are goroutines working and consuming
		select {
		case err := <-errors:
			fmt.Println("[ERROR]", err)
			drainChannel = true
			break MAINLOOP // Stop producing more stuff
		case inbox <- workToDo:
		}

	}

	close(inbox)
	if drainChannel {
		for range inbox {
			// Nothing, just consume the pending tasks
		}
	}
	// Wait until all the goroutines finish
	wg.Wait()

	return
}

func worker(workerID int, wg *sync.WaitGroup, inbox chan Workload, errors chan error) {
	defer wg.Done()

	// Listen for work from the inbox until the channel is closed
	for work := range inbox {
		endCounting := startCounting(workerID)

		if err := work.Process(); err != nil {
			errors <- fmt.Errorf("work.Process: %w", err)
			return
		}

		endCounting()
	}
}

func startCounting(workerID int) func() {
	start := time.Now()
	return func() {
		end := time.Now()
		_, _ = end, start
		//fmt.Printf("[%d] elapsed: %v \n", workerID, end.Sub(start))
	}
}

func main() {
	// Run the consumer on a couple test cases

	fmt.Println("test 1\n======")
	consume(2, testProducer(8))
	fmt.Println("")

	fmt.Println("test 2\n======")
	consume(4, testProducer(20))
	fmt.Println("")
}

//////////////////////////////////////////
// producer implementation details
// (these should not change)

type workload int

func (w workload) Process() error {
	fmt.Printf("Processing: %d\n", int(w))

	rand.Seed(time.Now().UnixMicro())
	if rand.Int()%5 == 0 {
		return errors.New("random error")
	}

	return nil
}

// testProducer creates an example producer with the given number
// of workloads.
func testProducer(n int) Producer {
	v := 0
	return func() (Workload, bool) {
		if v == n {
			return workload(0), true
		}

		i := v
		v++
		return workload(i), false
	}
}
