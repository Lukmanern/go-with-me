package main

import (
	"fmt"
	"strings"
)

// fuzzySearch returns a bool indicating whether the search string
// is found in the slice. It uses the strings.Contains() function
// to perform a fuzzy search.
func fuzzySearch(slice []string, search string) bool {
	for _, s := range slice {
		if strings.Contains(s, search) {
			return true
		}
	}
	return false
}

func main() {
	// Define a slice of strings to search in.
	slice := []string{"apple", "banana", "carrot", "daikon"}

	// Search for "an" in the slice. This should return true.
	found := fuzzySearch(slice, "an")
	fmt.Println(found)

	// Search for "z" in the slice. This should return false.
	found = fuzzySearch(slice, "z")
	fmt.Println(found)
}