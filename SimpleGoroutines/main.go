package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// This function generates a random number between 1 and 100.
// The function uses Go's concurrency features to perform the
// generation in a separate goroutine, and uses a channel to
// return the result to the calling function.
func generateRandomNumber() chan int {
	// Create a new channel to return the result.
	result := make(chan int)

	// Start a new goroutine to generate the random number.
	go func() {
		// Seed the random number generator with the current time.
		rand.Seed(time.Now().UnixNano())

		// Generate the random number and send it to the result channel.
		result <- rand.Intn(100) + 1
	}()

	// Return the result channel to the calling function.
	return result
}

func main() {
	// Use a wait group to wait for all goroutines to finish.
	var wg sync.WaitGroup

	// Generate 10 random numbers in separate goroutines.
	for i := 0; i < 10; i++ {
		// Increment the wait group counter.
		wg.Add(1)

		// Start a new goroutine to generate a random number.
		go func() {
			// Generate the random number.
			num := <-generateRandomNumber()

			// Print the random number.
			fmt.Println(num)

			// Decrement the wait group counter.
			wg.Done()
		}()
	}

	// Wait for all goroutines to finish.
	wg.Wait()
}