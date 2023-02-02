package main

import "fmt"

func main() {
	// Create a slice with 
	// duplicate elements
	s := []int{1, 2, 3, 3, 4, 5, 5}

	// Create a map to 
	// store unique elements
	m := make(map[int]bool)

	// Iterate over the 
	// slice with range
	for _, v := range s {
		// Add each element 
		// to the map with 
		// a value of true
		m[v] = true
	}

	// Create a slice to 
	// store the unique elements
	u := make([]int, 0, len(m))

	// Iterate over the map
	for k := range m {
		// Append each key (which is 
		// a unique element) to the slice
		u = append(u, k)
	}

	// Print the slice with unique elements
	fmt.Println("Unique elements:", u)
}