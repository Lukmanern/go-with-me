package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Define a struct that represents the data we want to
// extract from the JSON document.
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	// Define a JSON document as a string.
	jsonData := `
	{
		"name": "John Doe",
		"email": "johndoe@example.com",
		"password": "secret"
	}
	`
	// Use the json.Unmarshal function to parse the JSON data
	// and store the result in a User struct.
	var user User
	if err := json.Unmarshal([]byte(jsonData), &user); err != nil {
		log.Fatal(err)
	}
	// Print the user's name and email.
	fmt.Println("Name:", user.Name)
	fmt.Println("Email:", user.Email)
}