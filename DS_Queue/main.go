package main

import "fmt"

// Queue is a basic queue
// data structure.
type Queue []int

// Push adds a new element to 
// the end of the queue.
func (q *Queue) Push(i int) {
	*q = append(*q, i)
}

// Pop removes the first element 
// from the queue and returns it.
func (q *Queue) Pop() int {
	// Get the first element 
	// in the queue.
	x := (*q)[0]

	// Remove the first element 
	// from the queue.
	*q = (*q)[1:]

	return x
}

func main() {
	// Create a new queue.
	q := Queue{}

	// Push some elements 
	// onto the queue.
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)

	// Pop an element from the queue.
	fmt.Println(q.Pop()) // Output: 1

	// Pop another element from the queue.
	fmt.Println(q.Pop()) // Output: 2
}
