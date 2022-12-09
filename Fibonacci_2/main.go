package main

import "fmt"

func main() {
	Fib(1, 5) // 1, 1, 2, 3, 5,
	Fib(19, 5) // 19, 19, 38, 57, 95,
}


// start specifies the starting number in the sequence,
// and length specifies how many numbers should be generated. 
// For example, if start is 1 and length is 5, the function will 
// generate the first five numbers in the Fibonacci sequence (1, 1, 2, 3, 5).
// NOTE : start must more than 1 and length must more than 2.
func Fib(start uint, length uint) {
	if start == 0 || length <= 1 {
		return
	}

	var before uint = start
	Print(before)
	Print(start)

	for i := 3; i <= int(length); i++ {
		start = start + before
		before = start - before
		Print(start)
	}
	fmt.Println()
}

// Helper for print nth Fib numbers
func Print(n uint) {
	fmt.Print(n, ", ")
}