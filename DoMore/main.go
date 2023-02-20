package main

import "fmt"

func main() {
	doMore(someFunc1, 10)
}

func someFunc1() {
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
}

/* 
doMore is a recursive function that 
takes a function and a counter as input.
The function is executed once and 
the counter value is printed to the console.
The function is then called recursively 
with the same function and 
a decremented counter value.
The function returns 0 when the counter 
is less than or equal to 0.
*/
func doMore(fn func(), counter int) int {
	if counter <= 0 {
		// Return 0 to terminate/ kill/ stop 
		// the recursive call
		return 0
	}

	fn()
	fmt.Println("counter : ", counter)

	// Call doMore recursively with 
	// the same function and 
	// decremented counter value
	return doMore(fn, counter-1)
}