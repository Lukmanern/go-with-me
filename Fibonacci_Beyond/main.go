package main

import "fmt"

func main() {
	// Generate (calling) the Fibonacci sequence
	// and print it to the console.
	for n := range fibonacci() {
		fmt.Println(n)
		if n > 100 {
			break
		}
	}
}

// fibonacci is a function that returns a channel
// that generates the Fibonacci sequence.
func fibonacci() chan int {
	// new channel
	out := make(chan int)
	go func() {
		// Defer closing the channel to signal
		// to the caller that the sequence is complete
		defer close(out)
		for i, j := 0, 1; ; i, j = i+j, i {
			// Send the next number in
			// the sequence on the channel
			out <- i
			// Calculate the next two numbers in the sequence
			// i, j = i+j, i (its same, look up)
		}
	}()
	// Return the channel to the caller
	return out
}
