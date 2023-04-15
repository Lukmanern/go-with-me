package main

import (
	"fmt"
	"log"
)

func main() {
	// Call the divideByZero function
	divideByZero()
	// Print a message indicating
	// that we survived dividing by zero
	fmt.Println("we survived dividing by zero!")
}

// DivideByZero function to divide
// by zero and recover from panic
func divideByZero() {
	// Use defer to recover from
	// panic and log the error message
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	// Call the divide function
	// with arguments 1 and 0
	fmt.Println(divide(1, 0))
}

// Divide function to divide two integers
func divide(a, b int) int {
	// Check if the divisor
	// is 0, if so, panic
	if b == 0 {
		panic(nil)
	}
	// Return the result
	// of dividing a by b
	return a / b
}
