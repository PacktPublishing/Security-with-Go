package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"strconv"
)

func printUsage() {
	fmt.Printf(os.Args[0] + `

Generate a private and public RSA keypair and save as PEM files.
If no key size is provided, a default of 2048 is used.

Usage:
  ` + os.Args[0] + ` <private_key_filename> <public_key_filename> [keysize]

Examples:
  # Store generated private and public key in privkey.pem and pubkey.pem
  ` + os.Args[0] + ` priv.pem pub.pem
  ` + os.Args[0] + ` priv.pem pub.pem 4096`)
}

func checkArgs() (string, string, int) {
	// Too many or too few arguments
	if len(os.Args) < 3 || len(os.Args) > 4 {
		printUsage()
		os.Exit(1)
	}

	defaultKeySize := 2048

	// If there are 2 args provided, privkey and pubkey filenames
	if len(os.Args) == 3 {
		return os.Args[1], os.Args[2], defaultKeySize
	}

	// If 3 args provided, privkey, pubkey, keysize
	if len(os.Args) == 4 {
		keySize, err := strconv.Atoi(os.Args[3])
		if err != nil {
			printUsage()
			fmt.Println("Keysize not a valid number. Try 1024 or 2048.")
			os.Exit(1)
		}
		return os.Args[1], os.Args[2], keySize
	}

	return "", "", 0 // Default blank return catch-all
}

// Encode the private key as a PEM file
// PEM is a base-64 encoding of the key
func getPrivatePemFromKey(privateKey *rsa.PrivateKey) *pem.Block {
	encodedPrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	var privatePem = &pem.Block {
		Type: "RSA PRIVATE KEY",
		Bytes: encodedPrivateKey,
	}
	return privatePem
}

// Encode the public key as a PEM file
func generatePublicPemFromKey(publicKey rsa.PublicKey) *pem.Block {
	encodedPubKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		log.Fatal("Error marshaling PKIX pubkey. ", err)
	}

	// Create a public PEM structure with the data
	var publicPem = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: encodedPubKey,
	}
	return publicPem
}

func savePemToFile(pemBlock *pem.Block, filename string) {
	// Save public pem to file
	publicPemOutputFile, err := os.Create(filename)
	if err != nil {
		log.Fatal("Error opening pubkey output file. ", err)
	}
	defer publicPemOutputFile.Close()


	err = pem.Encode(publicPemOutputFile, pemBlock)
	if err != nil {
		log.Fatal("Error encoding public PEM. ", err)
	}
}

// Generate a public and private RSA key in PEM format
func main() {
	privatePemFilename, publicPemFilename, keySize := checkArgs()

	// Generate private key
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		log.Fatal("Error generating private key. ", err)
	}

	// Encode keys to PEM format
	privatePem := getPrivatePemFromKey(privateKey)
	publicPem := generatePublicPemFromKey(privateKey.PublicKey)

	// Save the PEM output to files
	savePemToFile(privatePem, privatePemFilename)
	savePemToFile(publicPem, publicPemFilename)

	// Print the public key to STDOUT for convenience
	fmt.Printf("%s", pem.EncodeToMemory(publicPem))
}
