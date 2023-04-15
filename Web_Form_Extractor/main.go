package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	// Get the URL of the page to parse
	url := "https://buildwithangga.com/search?keyword=golang"

	// Make a GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	// Convert the response body to a string
	html := string(body)

	// Find all HTML forms in the page
	forms := findForms(html)

	// Print the forms
	for _, form := range forms {
		fmt.Println(form)
	}
}

// findForms finds all HTML forms
// in the given string and returns
// a slice of strings containing the forms.
func findForms(html string) []string {
	var forms []string

	// Split the HTML string by the "<form"
	// string to get a slice of
	// strings containing all forms in the page
	formStrs := strings.Split(html, "<form")

	// Iterate over the slice of form
	// strings and extract the forms
	for _, formStr := range formStrs {
		if strings.Contains(formStr, "</form>") {
			// If the form string contains the "</form>"
			// string, then it is a valid
			// form, so add it to the slice of forms
			forms = append(forms, "<form"+formStr)
		}
	}

	return forms
}
