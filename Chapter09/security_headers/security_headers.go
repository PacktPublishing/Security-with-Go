package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
)

// Custom middleware handler logs user agent
func addSecureHeaders(rw http.ResponseWriter, r *http.Request,
	next http.HandlerFunc) {
	rw.Header().Add("Content-Security-Policy", "default-src 'self'")
	rw.Header().Add("X-Frame-Options", "SAMEORIGIN")
	rw.Header().Add("X-XSS-Protection", "1; mode=block")
	rw.Header().Add("Strict-Transport-Security", "max-age=10000, includeSubdomains; preload")
	rw.Header().Add("X-Content-Type-Options", "nosniff")

	next(rw, r) // Pass on to next middleware handler
}

// Return response to client
func indexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "You requested: "+request.URL.Path)
}

func main() {
	multiplexer := http.NewServeMux()
	multiplexer.HandleFunc("/", indexHandler)

	negroniHandler := negroni.New()

	// Set up as many middleware functions as you need, in order
	negroniHandler.Use(negroni.HandlerFunc(addSecureHeaders))
	negroniHandler.Use(negroni.NewLogger())
	negroniHandler.UseHandler(multiplexer)

	http.ListenAndServe("localhost:3000", negroniHandler)
}
