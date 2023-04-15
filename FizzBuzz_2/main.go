package main

import "fmt"

func main() {
	FizzBuzz(90)
}

// print 1 to n, replacing numbers that are
// divisible by 3 with "Fizz", numbers
// that are divisible by 5 with "Buzz",
// and numbers that are divisible by
// both 3 and 5 with "FizzBuzz".
// with triple if state
func FizzBuzz(n int) {
	if n < 3 {
		return
	}
	var s string
	for i := 1; i <= n; i++ {
		s = ""
		if i%3 == 0 {
			s += "Fizz"
		}
		if i%5 == 0 {
			s += "Buzz"
		}
		if s == "" {
			s = fmt.Sprint(i)
		}
		fmt.Println(s)
	}
}
