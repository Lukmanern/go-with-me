package main

import "fmt"

// factorial calculates the factorial of a given integer n
// using a recursive function.
func factorial(n int) int {
	// Print the value of n to the console.
	fmt.Println(n)
	
	// If n is 0, return 1 as the base case for the recursion.
	if n == 0 {
		return 1
	}
	
	// Otherwise, return n multiplied by the factorial of n-1.
	// This is the recursive step, where the function calls itself
	// with a simplified input (n-1) until it reaches the base case.
	return n * factorial(n-1)
}

func main() {
	// Call the factorial function with an input of 5.
	// This will calculate the factorial of 5, which is 5 x 4 x 3 x 2 x 1 = 120.
	// The result will be printed to the console.
	fmt.Println(factorial(5))
}
