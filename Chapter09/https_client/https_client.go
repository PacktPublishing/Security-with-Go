package main

import (
	"crypto/tls"
	"log"
	"net/http"
)

func main() {
	// Load cert
	cert, err := tls.LoadX509KeyPair("cert.pem", "privKey.pem")
	if err != nil {
		log.Fatal(err)
	}

	// Configure TLS client
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	// Use client to make request.
	// Ignoring response, just verifying connection accepted.
	_, err = client.Get("https://example.com")
	if err != nil {
		log.Println("Error making request. ", err)
	}
}
