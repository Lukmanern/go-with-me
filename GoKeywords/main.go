package main

import (
	"fmt"
	"time"
)

func main() {
	using_rare_keywords()
}

func using_rare_keywords() {
	var input int
	// Prompt user to enter a number
	fmt.Print("Enter a number: ")
	fmt.Scan(&input)

	// Check the input number
	switch input {
	case 1:
		fmt.Println("One")
		// fallthrough is used here to execute 
		// the next case statement but this is 
		// not recommended as it can lead to unexpected behavior
		fallthrough
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
		// goto statement is used here to jump 
		// to the end of the function but this is 
		// not recommended as it can make the code hard to follow
		goto end
	default:
		fmt.Println("Invalid input")
	}

	// select statement is used here to check for timeout
	select {
	case <- time.After(time.Second):
		fmt.Println("Timeout")
	}

end:
	// End of function message
	fmt.Println("End of function")
}

