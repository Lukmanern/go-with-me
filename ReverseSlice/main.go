package main

import (
	"fmt"
)

func main() {
	s := []int{4, 2, 7, 3, 8, 1}

	// Print the original slice
	fmt.Println("Original slice:", s)

	// Reverse the order of the slice
	// i is from the start => 0
	// j is from the back => len(s) - 1
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		// Swap
		s[i], s[j] = s[j], s[i]
	}

	// Print the reversed slice
	fmt.Println("Reversed slice:", s)
}
