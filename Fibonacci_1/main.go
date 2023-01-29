package main

import "fmt"

func main() {
	Fib(1, 5) // 1, 1, 2, 3, 5,
	Fib(19, 5) // 19, 19, 38, 57, 95,
}

// Helper function to print 
// a value followed by 
// a comma and space
func Print(n uint) {
	fmt.Print(n, ", ")
}

func Fib(start uint, length uint) {
	// Check if the starting value is 0 or 
	// if the length is less than or equal to 1
	// If either condition is true, 
	// return without printing anything
	if start == 0 || length <= 1 {
		return
	}

	// Store the starting value 
	// in a variable called "before"
	var before uint = start

	// Print the starting value
	Print(before)

	// Print the second value, 
	// which is the same
	//  as the starting value
	Print(start)

	// Loop through the remaining 
	// values up to the specified length
	for i := 3; i <= int(length); i++ {
		// Calculate the next Fibonacci 
		// number by adding the 
		// current value to 
		// the previous value
		start = start + before

		// Update the "before" variable to be 
		// the current value minus 
		// the previous value
		before = start - before

		// Print the current value
		Print(start)
	}

	// Print a newline at the end 
	// to separate the 
	// different sequences
	fmt.Println()
}
