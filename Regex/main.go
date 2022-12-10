package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Define a regex pattern for email matching
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)

	// Read an email from the standard input
	var email string
	fmt.Print("Enter an email address: ")
	fmt.Scanln(&email)

	// Check if the email matches the regex pattern
	if emailRegex.MatchString(email) {
		fmt.Println("The email is valid")
	} else {
		fmt.Println("The email is not valid")
	}

	// Define a regex pattern for phone number matching
	phoneRegex := regexp.MustCompile(`^[+]*[(]{0,1}[0-9]{1,4}[)]{0,1}[-\s\./0-9]*$`)

	// Read a phone number from the standard input
	var phoneNumber string
	fmt.Print("Enter a phone number: ")
	fmt.Scanln(&phoneNumber)

	// Check if the phone number matches the regex pattern
	if phoneRegex.MatchString(phoneNumber) {
		fmt.Println("The phone number is valid")
	} else {
		fmt.Println("The phone number is not valid")
	}

	// Define a regex pattern for URL matching
	urlRegex := regexp.MustCompile(`^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/)?[a-z0-9]+([-.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`)

	// Read a URL from the standard input
	var url string
	fmt.Print("Enter a URL: ")
	fmt.Scanln(&url)

	// Check if the URL matches the regex pattern
	if urlRegex.MatchString(url) {
		fmt.Println("The URL is valid")
	} else {
		fmt.Println("The URL is not valid")
	}
}
