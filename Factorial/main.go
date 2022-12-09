package main

import "fmt"

func factorial(n int) int {
	fmt.Println(n)
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(5))
}