package main

import "fmt"

func main() {
	fmt.Println(Pow(5, 5))
}

// This function calculates the
// value of x raised to the power of n.
func Pow(x float64, n int) float64 {
	// If x is equal to 1, return 1.
	if x == 1 {
		return 1
	}
	// If n is equal to 0, return 1.
	if n == 0 {
		return 1
	}
	// If n is negative, return the value
	// of (1/x) raised to the power of -n.
	if n < 0 {
		return Pow(1/x, -n)
	}
	// Calculate the value of
	// x raised to the power of n/2.
	half := Pow(x, n/2)
	// If n is even, return the
	// square of the value calculated above.
	if n%2 == 0 {
		return half * half
	}
	// If n is odd, return the value
	// calculated above multiplied by x.
	return half * half * x
}
