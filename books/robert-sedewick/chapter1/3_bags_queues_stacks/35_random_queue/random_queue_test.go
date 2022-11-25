package random_queue_test

import (
	"fmt"
	random_queue "github.com/jan-carreras/go-playground/books/robert-sedewick/chapter1/3_bags_queues_stacks/35_random_queue"
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
