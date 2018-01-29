package main

import (
	"fmt"
	"log"
	"net/http"
)

// Correctly formatted function declaration to satisfy the
// Go http.Handler interface. Any function that has the proper
// request/response parameters can be used to process an HTTP request.
// Inside the request struct we have access to the info about
// the HTTP request and the remote client.
func logRequest(response http.ResponseWriter, request *http.Request) {
	// Write output to file or just redirect output of this program to file
	log.Println(request.Method + " request from " + request.RemoteAddr + ". " +
		request.RequestURI)
	// If POST not empty, log attempt.
	username := request.PostFormValue("username")
	password := request.PostFormValue("pass")
	if username != "" || password != "" {
		log.Println("Username: " + username)
		log.Println("Password: " + password)
	}

	fmt.Fprint(response, "<html><body>")
	fmt.Fprint(response, "<h1>Login</h1>")
	if request.Method == http.MethodPost {
		fmt.Fprint(response, "<p>Invalid credentials.</p>")
	}
	fmt.Fprint(response, "<form method=\"POST\">")
	fmt.Fprint(response,
		"User:<input type=\"text\" name=\"username\"><br>")
	fmt.Fprint(response,
		"Pass:<input type=\"password\" name=\"pass\"><br>")
	fmt.Fprint(response, "<input type=\"submit\"></form><br>")
	fmt.Fprint(response, "</body></html>")
}

func main() {
	// Tell the default server multiplexer to map the landing URL to
	// a function called logRequest
	http.HandleFunc("/", logRequest)

	// Kick off the listener using that will run forever
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting listener. ", err)
	}
}
