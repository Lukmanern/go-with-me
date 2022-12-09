package main

import "fmt"

func main() {
	fmt.Println(Pow(5, 5))
}

// It uses recursion to repeatedly divide the 
// exponent by 2 until it reaches a value 
// of 0 or 1, at which point it can 
// return the result.
func Pow(x float64, n int) float64 {
	if x == 1 {
		return 1
	}
	if n == 0 {
		return 1
	}
	if n < 0 {
		return Pow(1/x, -n)
	}
	half := Pow(x, n/2)
	if n%2 == 0 {
		return half*half
	}
	return half*half*x
}