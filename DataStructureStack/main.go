package main

import "fmt"

// Stack is a basic stack data structure.
type Stack []int

// Push adds a new element to the top of the stack.
func (s *Stack) Push(i int) {
	*s = append(*s, i)
}

// Pop removes the top element from the stack and returns it.
func (s *Stack) Pop() int {
	// Get the length of the stack.
	n := len(*s)

	// Get the last element in the stack.
	x := (*s)[n-1]

	// Remove the last element from the stack.
	*s = (*s)[:n-1]

	return x
}

func main() {
	// Create a new stack.
	s := Stack{}

	// Push some elements onto the stack.
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	// Pop an element from the stack.
	fmt.Println(s.Pop()) // Output: 4

	// Pop another element from the stack.
	fmt.Println(s.Pop()) // Output: 3
}