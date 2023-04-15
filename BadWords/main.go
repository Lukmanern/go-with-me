package main

import (
	"fmt"
	"strings"
)

// list of bad words to be filtered
var BadWords []string = []string{"fuck", "damn", "crap"}

func main() {
	addBadWords("shit")
	input := "How the fuck are you doing that flip? shit, that was so cool men !"
	// call filterBadWords function to filter
	// the bad words from input string
	filteredInput := filterBadWords(input)
	fmt.Println(filteredInput)
}

// function to filter bad words from input string
func filterBadWords(input string) string {
	// iterate through each bad word in the list
	for _, word := range BadWords {
		// replace bad word with ****
		input = strings.ReplaceAll(input, word, "****")
	}
	return input
}

// function to add new bad words to the list
func addBadWords(words ...string) {
	BadWords = append(BadWords, words...)
}
