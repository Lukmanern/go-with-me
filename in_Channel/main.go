package main

import "fmt"

func main() {
	// Create a channel of type int
	ch := make(chan int)

	// Launch a goroutine that sends 
	// the value 42 to the channel
	go func() {
		ch <- 42
	}()

	// Use the <- syntax to receive 
	// a value from the channel
	fmt.Println(<-ch) // Output: 42
}