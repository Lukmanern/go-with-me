package main

import (
	"encoding/json"
	"net/http"
)

// User represents a user with a name and age
type User struct {
	Name string `json:"name"` 
	Age  int    `json:"age"`
}

func main() {
	// set up an HTTP handler function that handles all requests to the root path "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// create a new instance of the User struct
		user := User{Name: "John", Age: 30}

		// set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")
		// encode the user struct as JSON and write it to the response writer
		json.NewEncoder(w).Encode(user)
	})

	// start the HTTP server on port 8080 and handle requests with the handler function
	http.ListenAndServe(":8080", nil)
}
