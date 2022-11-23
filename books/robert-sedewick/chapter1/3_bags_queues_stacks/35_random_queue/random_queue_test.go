package random_queue_test

import (
	random_queue "exercises/books/robert-sedewick/chapter1/3_bags_queues_stacks/35_random_queue"
	"fmt"
	"testing"
)

func TestSample(t *testing.T) {

	r := new(random_queue.RandomQueue[string])
	r.Enqueue("hello")
	r.Enqueue("beautiful")
	r.Enqueue("world")

	// TODO: Assertions not done
	fmt.Println(r.Sample())
	fmt.Println(r.Sample())
	fmt.Println(r.Sample())
}
