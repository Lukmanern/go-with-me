package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	// Make an HTTP GET request to the 
	// website you want to extract URLs from
	resp, err := http.Get("https://buildwithangga.com/search?keyword=golang")
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return
	}

	// Read the response body and close the response
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading HTTP response body:", err)
		return
	}

	// Create a regular expression to match URLs
	urlRegex := regexp.MustCompile(`https?://[^\s"]+`)

	// Find all URLs that match the regular expression
	matches := urlRegex.FindAllString(string(body), -1)

	// Print the matches
	for _, match := range matches {
		fmt.Println(match)
	}
}
