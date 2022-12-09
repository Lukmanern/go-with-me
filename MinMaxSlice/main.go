package main

import "fmt"

func main() {
	s := []int{10, 2, -5, 4, 8, -3}
	fmt.Println(MinMaxSlice(s))
}

func MinMaxSlice(s []int) (int, int) {
	// Initialize the min and max value
	// to the first element of the slice
	min := s[0]
	max := s[0]

	// Iterate over the slice with range
	for _, v := range s {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}