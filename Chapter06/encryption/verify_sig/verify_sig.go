package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func printUsage() {
	fmt.Println(os.Args[0] + `

Verify an RSA signature of a message using SHA-256 hashing.
Public key is expected to be a PEM file.

Usage:

  ` + os.Args[0] + ` <publicKeyFilename> <signatureFilename> <messageFilename>

Example:
  ` + os.Args[0] + ` pubkey.pem signature.txt message.txt
`)
}

// Get arguments from command line
func checkArgs() (string, string, string) {
	// Expect 3 arguments: pubkey, signature, message file names
	if len(os.Args) != 4 {
		printUsage()
		os.Exit(1)
	}

	return os.Args[1], os.Args[2], os.Args[3]
}

// Returns bool whether signature was verified
func verifySignature(
	signature []byte,
	message []byte,
	publicKey *rsa.PublicKey) bool {

	hashedMessage := sha256.Sum256(message)

	err := rsa.VerifyPKCS1v15(
		publicKey,
		crypto.SHA256,
		hashedMessage[:],
		signature,
	)

	if err != nil {
		log.Println(err)
		return false
	}
	return true // If no error, match.
}

// Load file to memory
func loadFile(filename string) []byte {
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return fileData
}

// Load a public RSA key from a PEM encoded file
func loadPublicKeyFromPemFile(publicKeyFilename string) *rsa.PublicKey {
	// Quick load file to memory
	fileData, err := ioutil.ReadFile(publicKeyFilename)
	if err != nil {
		log.Fatal(err)
	}

	// Get the block data from the PEM encoded file
	block, _ := pem.Decode(fileData)
	if block == nil || block.Type != "PUBLIC KEY" {
		log.Fatal("Unable to load valid public key. ")
	}

	// Parse the bytes and store in a public key format
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatal("Error loading public key. ", err)
	}

	return publicKey.(*rsa.PublicKey) // Cast interface to PublicKey
}

// Verify a cryptographic signature using RSA PKCS#1 v1.5 with SHA-256
// and a PEM encoded PKIX public key.
func main() {
	// Parse command line arguments
	publicKeyFilename, signatureFilename, messageFilename :=
		checkArgs()

	// Load all the files from disk
	publicKey := loadPublicKeyFromPemFile(publicKeyFilename)
	signature := loadFile(signatureFilename)
	message := loadFile(messageFilename)

	// Verify signature
	valid := verifySignature(signature, message, publicKey)

	if valid {
		fmt.Println("Signature verified.")
	} else {
		fmt.Println("Signature could not be verified.")
	}
}
