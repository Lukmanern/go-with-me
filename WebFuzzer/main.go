package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Initialize an escapeList
	// slice with a list of strings
	escapeList := []string{
		"%00",
		"%0a",
		"%0a%20",
		"%0d",
		"\r\t",
	}

	// Set the target URL
	target := "http://example.xyz/?data="

	// Loop over the escapeList slice
	for _, escape := range escapeList {
		// Concatenate the escape string
		// with the target URL
		url := target + escape

		// Create a new GET request using the url
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			// If there is an error creating the request,
			// log an error and continue
			// to the next iteration
			fmt.Printf("[ERR] %s\n ::1", url)
			continue
		}

		// Create a new HTTP client
		client := &http.Client{}

		// Execute the request
		resp, err := client.Do(req)
		if err != nil {
			// If there is an error executing the request,
			// log an error and continue
			// to the next iteration
			fmt.Printf("[ERR] %s\n ::2", url)
			continue
		}

		// Defer closing the response body
		defer resp.Body.Close()

		// Check the status code of the response
		if resp.StatusCode != 200 {
			// If the status code is not 200, log an error
			fmt.Printf("[ERR] %s\n ::3", url)
		} else {
			// If the status code is 200, log success
			fmt.Printf("[OK] %s\n", url)
		}
	}
}
