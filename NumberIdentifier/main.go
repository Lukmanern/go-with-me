package main

import (
	"fmt"
	"math"
)

func main() {
	num := 5 // Example number

	// Check if the number is positive
	if num > 0 {
		fmt.Println("The number is positive.")
	}

	// Check if the number is odd
	if num%2 != 0 {
		fmt.Println("The number is odd.")
	}

	// Check if the number is prime
	isPrime := true
	maxFactor := int(math.Sqrt(float64(num)))
	for i := 2; i <= maxFactor; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println("The number is prime.")
	}
}
