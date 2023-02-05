package main

import (
	"fmt"
	"math"
)

func main() {
	var num int = 113 // Example number
	posOrNeg(num)
	oddOrEven(num)
	isPrime(num)
}

func posOrNeg(num int) {
	// Check if the number is positive
	if num >= 0 {
		fmt.Println("The number is positive.")
		// return for stop the func
		return
	}

	fmt.Println("The number is negative.")
}

func oddOrEven(num int) {
	if num < 0 {
		fmt.Println("The number is negative, cant check odd or even.")
		return
	}
	// Check if the number is odd
	if num%2 == 0 {
		fmt.Println("The number is even.")
		// return for stop the func
		return
	}

	fmt.Println("The number is odd.")
}

func isPrime(num int) {
	// Check if the number is less than 2, which is not a prime number
	if num < 2 {
		fmt.Println("The number is not prime. It's lower than 2.")
		return
	}

	// Find the maximum factor to check for the number being prime
	maxFactor := int(math.Sqrt(float64(num)))

	// Loop through all possible factors from 2 to the maxFactor
	for i := 2; i <= maxFactor; i++ {
		// Check if the number can be divided by the current factor
		if num%i == 0 {
			fmt.Printf("The number is not prime. (can divide by %d)\n", i)
			return
		}
	}

	// If the number cannot be divided by any of the factors, it's a prime number
	fmt.Println("The number is prime.")
}


