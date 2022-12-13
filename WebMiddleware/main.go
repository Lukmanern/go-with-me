package main

import (
	"fmt"
	"net/http"
)

func main() {
	// with middleware
	http.HandleFunc("/", logMiddleware(indexHandler))
	// without middleware
	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "> Home Page")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "> About Page")
}

// Middleware function to log the path of each incoming request
func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		next(w, r)
	}
}
