package main

import "fmt"

// isPalindrome returns true if the given slice is a palindrome, false otherwise
func isPalindrome(s []int) bool {
	// Compare the elements at the start and end of the slice, moving inward
	// Stop when the pointers meet in the middle of the slice
	// i from front, j from back
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			// The slice is not a palindrome
			return false
		}
	}

	// The slice is a palindrome
	return true
}

func main() {
	// Create a slice
	s := []int{1, 2, 3, 4, 3, 2, 1}

	// Check if the slice is a palindrome
	if isPalindrome(s) {
		fmt.Println("The slice is a palindrome")
	} else {
		fmt.Println("The slice is not a palindrome")
	}
}