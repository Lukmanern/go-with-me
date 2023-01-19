package main

import (
	"fmt"
	"math"
)

func main() {
	var num int = -113 // Example number
	pos_or_neg(num)
	odd_or_even(num)
	is_prime(num)
}

func pos_or_neg(num int) {
	// Check if the number is positive
	if num >= 0 {
		fmt.Println("The number is positive.")
		return
	}

	fmt.Println("The number is negative.")
}

func odd_or_even(num int) {
	// Check if the number is odd
	if num%2 != 0 {
		fmt.Println("The number is odd.")
		return
	}

	fmt.Println("The number is even.")
}

func is_prime(num int) {
	// Check if the number is prime
	if num < 2 {
		fmt.Printf("The number is not prime. It's lower than 2.")
		return
	}
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
		return
	}
	fmt.Printf("The number is not prime. (can divide by %d)\n", i)
}
