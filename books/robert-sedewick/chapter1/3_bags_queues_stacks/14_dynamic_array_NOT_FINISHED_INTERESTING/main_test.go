package main

import (
	"fmt"
	"testing"
)

func TestEnqueue(t *testing.T) {
	q := new(Queue)
	q.Enqueue("hello1")
	q.Enqueue("hello2")
	q.Enqueue("hello3")
	q.Enqueue("hello4")
	q.Enqueue("hello5")
	q.Enqueue("hello6")
	q.Enqueue("hello7")
	q.Enqueue("hello8")
	q.Enqueue("hello9")
	q.Enqueue("hello10")
	fmt.Println(q.Dequeue())
}
