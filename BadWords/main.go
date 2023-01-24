package main

import (
	"fmt"
	"strings"
)

var BadWords []string = []string{"fuck", "damn", "crap"}

func main() {
	input := "How the fuck are you doing that flip?"
	filteredInput := filterBadWords(input)
	fmt.Println(filteredInput)
}

func filterBadWords(input string) string {
	for _, word := range BadWords {
		input = strings.ReplaceAll(input, word, "****")
	}
	return input
}

func addBadWords(words... string) {
	BadWords = append(BadWords, words...)
}
