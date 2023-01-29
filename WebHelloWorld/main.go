package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Handle requests to the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Write the "Hello, world!" 
		// message to the response
		logging(r.RemoteAddr, "{__blank__}");
		fmt.Fprintln(w, "Hello, world!")
	})

	http.HandleFunc("/example", func(w http.ResponseWriter, r *http.Request) {
		logging(r.RemoteAddr, "example");
		fmt.Fprintln(w, "This Is Another Shot !")
	})

	// geting all url-path that user give to server
	http.HandleFunc("/read/", func(w http.ResponseWriter, r *http.Request) {
		logging(r.RemoteAddr, "read-path");
		// Get the value of the "path"
		path := r.URL.Path
		// Print the value of the "path"
		fmt.Fprintln(w, "Path :", path)
	})

	// Start the server on port 8080
	fmt.Println("> http://localhost:8080/")
	fmt.Println("> http://localhost:8080/example")
	fmt.Println("> http://localhost:8080/read/what_ever/have_fun")
	// you can delete 0.0.0.0 for localhost
	http.ListenAndServe("0.0.0.0:8080", nil)
}

// this is simple logging, 
// read IP Address and URL parameter
func logging(ip string, param string) {
	log.Printf("> ip:%s parameter:%s\n", ip, param)
}