package main

import "fmt"

func main() {
	sl := []int{
		1, 2, 3, 4, 5,
	}
	fmt.Println(sl) // 1, 2, 3, 4, 5,

	doubleTheSlice(sl)
	fmt.Println(sl) // 2, 4, 6, 8, 10,
}

// This func doesn't have
// return value
func doubleTheSlice(s []int) {
	for i := range s {
		// double the value
		s[i] *= 2
	}
}