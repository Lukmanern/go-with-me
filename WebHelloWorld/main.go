package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Handle requests to the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Write the "Hello, world!" message to the response
		fmt.Fprintln(w, "Hello, world!")
	})

	http.HandleFunc("/example", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This Is Another Shot !")
	})

	// Start the server on port 8080
	fmt.Println("> http://localhost:8080/")
	fmt.Println("> http://localhost:8080/example")
	http.ListenAndServe(":8080", nil)
}