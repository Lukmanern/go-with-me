package main

import "fmt"

func main() {
	FibPrinter(20)
}

// FibPrinter loop n-times and print the n result.
func FibPrinter(n int) {
	for i := 1; i <= n; i++ {
		fmt.Print(Fib(i), ", ")
	}
}

// Fib returns the n-th Fibonacci number.
func Fib(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}