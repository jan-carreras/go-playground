package main

import "fmt"

// 1.3.14 Develop a class ResizingArrayQueueOfStrings that implements the queue
// abstraction with a fixed-size array, and then extend your implementation to
// use array resizing to remove the size restriction.

// Personal note: I should use slices, but I cannot call "append"

type Queue struct {
	arr      []string
	length   int
	capacity int
}

func (q *Queue) Enqueue(s string) {
	if q.atCapacity() {
		q.resize(q.length * 2)
	}
	q.arr[q.length] = s
	q.length++
	fmt.Println(">", q.length, q.capacity, q.arr)
}

func (q *Queue) Dequeue() string {
	if q.length == 0 {
		panic("nothing else to return")
	}
	res := q.arr[q.length]
	q.length--

	fmt.Println("<", q.length, q.capacity, q.arr)
	return res
}

func (q *Queue) atCapacity() bool {
	return q.capacity == q.length
}

func (q *Queue) resize(newSize int) {
	if newSize == 0 {
		newSize = 1
	}

	q.capacity = newSize

	fmt.Println("resized to... ", newSize)
	resized := make([]string, newSize)
	copy(resized, q.arr)
	q.arr = resized
}
