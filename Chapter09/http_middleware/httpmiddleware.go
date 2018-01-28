package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/urfave/negroni"
)

// Custom middleware handler logs user agent
func customMiddlewareHandler(rw http.ResponseWriter,
	r *http.Request,
	next http.HandlerFunc,
) {
	log.Println("Incoming request: " + r.URL.Path)
	log.Println("User agent: " + r.UserAgent())

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
	negroniHandler.Use(negroni.HandlerFunc(customMiddlewareHandler))
	negroniHandler.UseHandler(multiplexer)

	http.ListenAndServe("localhost:3000", negroniHandler)
}
