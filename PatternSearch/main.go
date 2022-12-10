package main

import (
	"fmt"
	"regexp"
)

func main() {
	// The pattern to search for
	pattern := "foo"
	// The text to search in
	text := "My text contains a pattern: foo"

	// Compile the regular expression
	regex, err := regexp.Compile(pattern)
	if err != nil {
		// Handle error
		fmt.Println("Error compiling regular expression:", err)
		return
	}

	// Search for the pattern in the text
	match := regex.MatchString(text)
	if match {
		fmt.Println("Pattern found in text")
	} else {
		fmt.Println("Pattern not found in text")
	}
}
