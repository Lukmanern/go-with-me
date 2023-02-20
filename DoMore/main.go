package main

import "fmt"

func main() {
	doMore(someFunc1, 10)
}

func someFunc1() {
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
}

func doMore(fn func(), counter int) int {
	if counter <= 0 {
		return 0
	}

	fn()
	fmt.Println("counter : ", counter)

	return doMore(fn, counter-1)
}