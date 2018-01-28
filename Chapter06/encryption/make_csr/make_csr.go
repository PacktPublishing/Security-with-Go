package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func printUsage() {
	fmt.Println(os.Args[0] + ` - Create a certificate signing request with a private key

Private key is expected in PEM format. Certificate valid for localhost only.
Certificate signing request is created using the SHA-256 hash.

Usage:
  ` + os.Args[0] + ` <privateKeyFilename> <csrOutputFilename>

Example:
  ` + os.Args[0] + ` priv.pem csr.pem
`)
}

func checkArgs() (string, string) {
	if len(os.Args) != 3 {
		printUsage()
		os.Exit(1)
	}

	// Private key filename, cert signing request output filename
	return os.Args[1], os.Args[2]
}

// Load the RSA private key from a PEM encoded file
func loadPrivateKeyFromPemFile(privateKeyFilename string) *rsa.PrivateKey {
	// Quick load file to memory
	fileData, err := ioutil.ReadFile(privateKeyFilename)
	if err != nil {
		log.Fatal("Error loading private key file. ", err)
	}

	// Get the block data from the PEM encoded file
	block, _ := pem.Decode(fileData)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		log.Fatal("Unable to load a valid private key.")
	}

	// Parse the bytes and put it in to a proper privateKey struct
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Fatal("Error loading private key.", err)
	}

	return privateKey
}

// Create a CSR PEM and save to file
func saveCSRToPemFile(csr []byte, filename string) {
	csrPem := &pem.Block{
		Type:  "CERTIFICATE REQUEST",
		Bytes: csr,
	}
	csrOutfile, err := os.Create(filename)
	if err != nil {
		log.Fatal("Error opening "+filename+" for saving. ", err)
	}
	pem.Encode(csrOutfile, csrPem)
}

// Create a certificate signing request with a private key valid for localhost
func main() {
	// Load parameters
	privKeyFilename, csrOutFilename := checkArgs()
	privKey := loadPrivateKeyFromPemFile(privKeyFilename)

	// Prepare information about organization the cert will belong to
	nameInfo := pkix.Name{
		Organization:       []string{"My Organization Name"},
		CommonName:         "localhost",
		OrganizationalUnit: []string{"Business Unit Name"},
		Country:            []string{"US"}, // 2-character ISO code
		Province:           []string{"Texas"},
		Locality:           []string{"Houston"}, // City
	}

	// Prepare CSR template
	csrTemplate := x509.CertificateRequest{
		Version:            2, // Version 3, zero-indexed values
		SignatureAlgorithm: x509.SHA256WithRSA,
		PublicKeyAlgorithm: x509.RSA,
		PublicKey:          privKey.PublicKey,
		Subject:            nameInfo,

		// Subject Alternate Name values.
		DNSNames:       []string{"Business Unit Name"},
		EmailAddresses: []string{"test@localhost"},
		IPAddresses:    []net.IP{},
	}

	// Create the CSR based off the template
	csr, err := x509.CreateCertificateRequest(rand.Reader, &csrTemplate, privKey)
	if err != nil {
		log.Fatal("Error creating certificate signing request. ", err)
	}
	saveCSRToPemFile(csr, csrOutFilename)
}
