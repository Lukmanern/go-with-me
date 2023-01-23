# Web Middleware

This is a simple Go program that creates an HTTP server and handles two routes: `/` and `/about`.

The main function sets up two http.HandleFunc handlers. The first handler, `indexHandler`, is wrapped in a middleware function called logMiddleware which will log the path of each incoming request. The second handler, aboutHandler, is not wrapped in middleware.

The `indexHandler` and `aboutHandler` functions both take in an http.ResponseWriter and an `http.Request` as arguments and writes a simple string to the ResponseWriter indicating which page the user is visiting.

The `logMiddleware` function is a middleware function that takes in a http.HandlerFunc and returns a new http.HandlerFunc which logs the path of the incoming request and calls the original handler function. The program then starts the HTTP server on port 8080 and handles requests with the handler functions.
