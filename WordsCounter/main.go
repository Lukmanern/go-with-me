package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read the input string
	fmt.Print("Enter a string: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Create a map to store 
	// the counts of each word
	wordCounts := make(map[string]int)

	// Split the input string into words
	words := strings.Fields(input)

	// Iterate through the words and
	//  update the count for each word
	for _, word := range words {
		wordCounts[word]++
	}

	// Print the count for each word
	for word, count := range wordCounts {
		fmt.Printf("%s: %d\n", word, count)
	}
}
