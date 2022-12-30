package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Struct to hold a single cat fact
type CatFact struct {
    	Fact string `json:"fact"`
}

func main() {
	// Make an HTTP GET request to the API
	resp, err := http.Get("https://catfact.ninja/fact")
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

	// Unmarshal the response body into a CatFact struct
	var fact CatFact
	err = json.Unmarshal(body, &fact)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Print the cat fact
	fmt.Println("Cat Fact:", fact.Fact)
}
