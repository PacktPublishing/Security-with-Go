package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509/pkix"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net"
	"os"
	"time"
)

func printUsage() {
	fmt.Println(os.Args[0] + ` - Generate a self signed TLS certificate

Usage:
  ` + os.Args[0] + ` <privateKeyFilename> <certOutputFilename> [-ca|--cert-authority]

Example:
  ` + os.Args[0] + ` priv.pem cert.pem
  ` + os.Args[0] + ` priv.pem cacert.pem -ca
`)
}

func checkArgs() (string, string, bool) {
	if len(os.Args) < 3 || len(os.Args) > 4 {
		printUsage()
		os.Exit(1)
	}

	// See if the last cert authority option was passed
	isCA := false // Default
	if len(os.Args) == 4 {
		if os.Args[3] == "-ca" || os.Args[3] == "--cert-authority" {
			isCA = true
		}
	}

	// Private key filename, cert output filename, is cert authority
	return os.Args[1], os.Args[2], isCA
}

func setupCertificateTemplate(isCA bool) x509.Certificate {
	// Set valid time frame to start now and end one year from now
	notBefore := time.Now()
	notAfter := notBefore.Add(time.Hour * 24 * 365) // 1 year/365 days

	// Generate secure random serial number
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	randomNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		log.Fatal("Error generating random serial number. ", err)
	}

	nameInfo := pkix.Name{
		Organization: []string{"My Organization"},
		CommonName: "localhost",
		OrganizationalUnit: []string{"My Business Unit"},
		Country:            []string{"US"}, // 2-character ISO code
		Province:           []string{"Texas"}, // State
		Locality:           []string{"Houston"}, // City
	}

	// Create the certificate template
	certTemplate := x509.Certificate{
		SerialNumber: randomNumber,
		Subject: nameInfo,
		EmailAddresses: []string{"test@localhost"},
		NotBefore: notBefore,
		NotAfter: notAfter,
		KeyUsage: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		// For ExtKeyUsage, default to any, but can specify to use
		// only as server or client authentication, code signing, etc
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageAny },
		BasicConstraintsValid: true,
		IsCA: false,
	}

	// To create a certificate authority that can sign cert signing requests, set these
	if isCA {
	    certTemplate.IsCA = true
	    certTemplate.KeyUsage = certTemplate.KeyUsage | x509.KeyUsageCertSign
	}

	// Add any IP addresses and hostnames covered by this cert
	// This example only covers localhost
	certTemplate.IPAddresses = []net.IP{net.ParseIP("127.0.0.1")}
	certTemplate.DNSNames = []string{"localhost", "localhost.local"}

	return certTemplate
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
		log.Fatal("Error loading private key. ", err)
	}

	return privateKey
}

// Save the certificate as a PEM encoded file
func writeCertToPemFile(outputFilename string, derBytes []byte ) {
	// Create a PEM from the certificate
	certPem := &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}

	// Open file for writing
	certOutfile, err := os.Create(outputFilename)
	if err != nil {
		log.Fatal("Unable to open certificate output file. ", err)
	}
	pem.Encode(certOutfile, certPem)
	certOutfile.Close()
}

// Create a self-signed TLS/SSL certificate for localhost with an RSA private key
func main() {
	privPemFilename, certOutputFilename, isCA := checkArgs()

	// Private key of signer - self signed means signer==signee
	privKey := loadPrivateKeyFromPemFile(privPemFilename)

	// Public key of signee. Self signing means we are the signer and the signee
	// so we can just pull our public key from our private key
	pubKey := privKey.PublicKey

	// Set up all the certificate info
	certTemplate := setupCertificateTemplate(isCA)

	// Create (and sign with the priv key) the certificate
	certificate, err := x509.CreateCertificate(
		rand.Reader,
		&certTemplate,
		&certTemplate,
		&pubKey,
		privKey,
	)
	if err != nil {
		log.Fatal("Failed to create certificate. ", err)
	}

	// Format the certificate as a PEM and write to file
	writeCertToPemFile(certOutputFilename, certificate)
}

