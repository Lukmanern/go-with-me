package main

import (
	"fmt"
	"math"
)

func main() {
	var num int = 112 // Example number

	// Check if the number is positive
	if num >= 0 {
		fmt.Println("The number is positive.")
	} else {
		fmt.Println("The number is negative.")
	}

	// Check if the number is odd
	if num%2 != 0 {
		fmt.Println("The number is odd.")
	} else {
		fmt.Println("The number is even.")
	}

	// Check if the number is prime
	var i		  int
	var isPrime   bool = true
	var maxFactor int  = int(math.Sqrt(float64(num)))
	for i = 2; i <= maxFactor; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println("The number is prime.")
	} else {
		fmt.Printf("The number is not prime. (can divide by %d)\n", i)
	}
}
