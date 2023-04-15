package main

import "fmt"

func main() {
	// Call the FibPrinter function with an input of 20.
	// This will print the first 20 Fibonacci numbers.
	FibPrinter(20)
}

// FibPrinter loop n-times and print the n result.
func FibPrinter(n int) {
	// Loop n-times, starting from 1.
	for i := 1; i <= n; i++ {
		// Print the n-th Fibonacci number
		// followed by a comma and a space.
		fmt.Print(Fib(i), ", ")
	}
}

// Fib returns the n-th Fibonacci number.
func Fib(n int) int {
	// If n is less than 0, return 0.
	if n < 0 {
		return 0
	}

	// If n is 0 or 1, return n.
	if n == 0 || n == 1 {
		return n
	}

	// Otherwise, return the sum of the previous two
	// Fibonacci numbers (Fib(n-1) + Fib(n-2)).
	return Fib(n-1) + Fib(n-2)
}
