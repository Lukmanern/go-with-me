package main

import "fmt"

func main() {
	// Create a channel of type int
	ch := make(chan int)

	// Launch a goroutine that sends the value 42 to the channel
	go func() {
		ch <- 42
	}()

	// Use the <- syntax to receive a value from the channel
	fmt.Println(<-ch) // Output: 42
}

/*
In this example, we create a channel ch of type int using the make function.
Then, we launch a goroutine that sends the value 42 to the channel using the <- syntax.
Finally, we receive the value from the channel using the <- syntax and print it to the console.
Channels are a powerful tool for communication between goroutines and
can be used in a variety of ways to synchronize and coordinate their execution.
I hope this helps! Let me know if you have any other questions.
*/