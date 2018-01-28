package main

import (
	"fmt"
	"net/http"
	"log"
	"time"
)

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	secureSessionCookie := http.Cookie {
		Name: "SessionID",
		Value: "<secure32ByteToken>",
		Domain: "yourdomain.com",
		Path: "/",
		Expires: time.Now().Add(60 * time.Minute),
		HttpOnly: true, // Prevents JavaScript from accessing
		Secure: true, // Requires HTTPS
	}

	// Write cookie header to response
	http.SetCookie(writer, &secureSessionCookie)
	fmt.Fprintln(writer, "Cookie has been set.")
}



func main() {
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("Error creating server. ", err)
	}
}
