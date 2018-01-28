package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
)

// Return response to client
func indexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "You requested: " + request.URL.Path)
}

func main() {
	multiplexer := http.NewServeMux()
	multiplexer.HandleFunc("/", indexHandler)

	negroniHandler := negroni.New()
	negroniHandler.Use(negroni.NewLogger()) // Use Negroni's default logging middleware
	negroniHandler.UseHandler(multiplexer)

	http.ListenAndServe("localhost:3000", negroniHandler)
}
