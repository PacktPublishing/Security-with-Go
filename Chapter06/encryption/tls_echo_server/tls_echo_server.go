package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"os"
)

func printUsage() {
	fmt.Println(os.Args[0] + ` - Start a TLS echo server

Server will echo one message received back to client.
Provide a certificate and private key file in PEM format.
Host string in the format: hostname:port

Usage:
  ` + os.Args[0] + ` <certFilename> <privateKeyFilename> <hostString>

Example:
  ` + os.Args[0] + ` cert.pem priv.pem localhost:9999
`)
}

func checkArgs() (string, string, string) {
	if len(os.Args) != 4 {
		printUsage()
		os.Exit(1)
	}

	return os.Args[1], os.Args[2], os.Args[3]
}

// Create a TLS listener and echo back data received by clients.
func main() {
	certFilename, privKeyFilename, hostString := checkArgs()

	// Load the certificate and private key
	serverCert, err := tls.LoadX509KeyPair(certFilename, privKeyFilename)
	if err != nil {
		log.Fatal("Error loading certificate and private key. ", err)
	}

	// Set up certificates, host/ip, and port
	config := &tls.Config{
		// Specify server certificate
		Certificates: []tls.Certificate{serverCert},

		// By default no client certificate is required.
		// To require and validate client certificates, specify the
		// ClientAuthType to be one of:
		//     NoClientCert, RequestClientCert, RequireAnyClientCert,
		//     VerifyClientCertIfGiven, RequireAndVerifyClientCert)

		// ClientAuth: tls.RequireAndVerifyClientCert

		// Define the list of certificates you will accept as
		// trusted certificate authorities with ClientCAs.

		// ClientCAs: *x509.CertPool
	}

	// Create the TLS socket listener
	listener, err := tls.Listen("tcp", hostString, config)
	if err != nil {
		log.Fatal("Error starting TLS listener. ", err)
	}
	defer listener.Close()

	// Listen forever for connections
	for {
		clientConnection, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting client connection. ", err)
			continue
		}
		// Launch a goroutine(thread) to handle each connection
		go handleConnection(clientConnection)
	}
}

// Function that gets launched in a goroutine to handle client connection
func handleConnection(clientConnection net.Conn) {
	defer clientConnection.Close()
	socketReader := bufio.NewReader(clientConnection)
	for {
		// Read a message from the client
		message, err := socketReader.ReadString('\n')
		if err != nil {
			log.Println("Error reading from client socket. ", err)
			return
		}
		fmt.Println(message)

		// Echo back the data to the client.
		numBytesWritten, err := clientConnection.Write([]byte(message))
		if err != nil {
			log.Println("Error writing data to client socket. ", err)
			return
		}
		fmt.Printf("Wrote %d bytes back to client.\n", numBytesWritten)
	}
}
