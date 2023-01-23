# Web Endpoint

This is a simple Go program that serves a JSON representation of a user on an HTTP server. The user is hardcoded with the name "John" and age 30.

The program uses the `net/http` package to set up an `HTTP handler function` that handles all requests to the root path "/". The function creates a new instance of the User struct, sets the `Content-Type header` to `application/json`, and encodes the User struct as JSON and writes it to the response writer. The program then starts the HTTP server on port 8080 and handles requests with the handler function.
