package main

import "fmt"

func main() {
	FizzBuzz(90)
}

// print 1 to n, replacing numbers that are
// divisible by 3 with "Fizz", numbers
// that are divisible by 5 with "Buzz",
// and numbers that are divisible by
// both 3 and 5 with "FizzBuzz"
func FizzBuzz(n int) {
	if n < 3 {
		return
	}

	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
