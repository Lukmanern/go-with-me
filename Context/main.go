package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// set time for do cancelFunc()
	times := 20*time.Second

	// Create context with timeout and value
	ctx, cancelFunc := context.WithTimeout(context.Background(), times)
	defer cancelFunc()

	// running worker with context
	// this function will cancel because of limitation time
	// or because time.Sleep()
	go worker(ctx, "Hello, World!")

	// this can more longer or shorter from time
	// just change it for deep understanding
	time.Sleep(15 * time.Second)

	fmt.Println("Canceling context...")
}

func worker(ctx context.Context, request string) {
	// Waiting for cancelFunc
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			// worker here
			fmt.Println(request)
			time.Sleep(1 * time.Second)
		}
	}
}
