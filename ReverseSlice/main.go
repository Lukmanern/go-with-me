package main

import (
	"fmt"
)

func main() {
	s := []int{4, 2, 7, 3, 8, 1}

	// Print the original slice
	fmt.Println("Original slice:", s)
	reverseSlice(s)

	// Print the reversed slice
	fmt.Println("Reversed slice:", s)
}

// no need return value
// see https://github.com/Lukmanern/go-with-me/tree/master/PassedByValue
func reverseSlice(s []int) {
	// Reverse the order of the slice
	// i is from the start => 0
	// j is from the back => len(s) - 1
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		// Swap
		s[i], s[j] = s[j], s[i]
	}
}
