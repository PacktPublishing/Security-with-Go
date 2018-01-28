package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func printUsage() {
	fmt.Printf(os.Args[0] + `

Encrypt or decrypt a file using AES with a 256-bit key file.
This program can also generate 256-bit keys.

Usage:
  ` + os.Args[0] + ` [-h|--help]
  ` + os.Args[0] + ` [-g|--genkey]
  ` + os.Args[0] + ` <keyFile> <file> [-d|--decrypt]

Examples:
  # Generate a 32-byte (256-bit) key
  ` + os.Args[0] + ` --genkey

  # Encrypt with secret key. Output to STDOUT
  ` + os.Args[0] + ` --genkey > secret.key

  # Encrypt message using secret key. Output to ciphertext.dat
  ` + os.Args[0] + ` secret.key message.txt > ciphertext.dat

  # Decrypt message using secret key. Output to STDOUT
  ` + os.Args[0] + ` secret.key ciphertext.dat -d

  # Decrypt message using secret key. Output to message.txt
  ` + os.Args[0] + ` secret.key ciphertext.dat -d > cleartext.txt
`)
}

// Check command-line arguments.
// If the help or generate key functions are chosen
// they are run and then the program exits
// otherwise it returns  keyFile, file, decryptFlag.
func checkArgs() (string, string, bool) {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		printUsage()
		os.Exit(1)
	}

	// One arg provided
	if len(os.Args) == 2 {
		// Only -h, --help and --genkey are valid one-argument uses
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			printUsage() // Print help text
			os.Exit(0)   // Exit gracefully no error
		}
		if os.Args[1] == "-g" || os.Args[1] == "--genkey" {
			// Generate a key and print to STDOUT
			// User should redirect output to a file if needed
			key := generateKey()
			fmt.Printf(string(key[:])) // No newline
			os.Exit(0)                 // Exit gracefully

		}
	}

	// The only use options left is
	// encrypt <keyFile> <file> [-d|--decrypt]
	// If there are only 2 args provided, they must be the
	// keyFile and file without a decrypt flag.
	if len(os.Args) == 3 {
		return os.Args[1], os.Args[2], false // keyFile, file, decryptFlag
	}
	// If 3 args are provided, check that the last one is -d or --decrypt
	if len(os.Args) == 4 {
		if os.Args[3] != "-d" && os.Args[3] != "--decrypt" {
			fmt.Println("Error: Unknown usage.")
			printUsage()
			os.Exit(1) // Exit with error code
		}
		return os.Args[1], os.Args[2], true
	}

	return "", "", false // Default blank return
}

func generateKey() []byte {
	randomBytes := make([]byte, 32) // 32 bytes, 256 bit
	numBytesRead, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal("Error generating random key.", err)
	}
	if numBytesRead != 32 {
		log.Fatal("Error generating 32 random bytes for key.")
	}
	return randomBytes
}

// AES encryption
func encrypt(key, message []byte) ([]byte, error) {
	// Initialize block cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Create the byte slice that will hold encrypted message
	cipherText := make([]byte, aes.BlockSize+len(message))

	// Generate the Initialization Vector (IV) nonce
	// which is stored at the beginning of the byte slice
	// The IV is the same length as the AES blocksize
	iv := cipherText[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return nil, err
	}

	// Choose the block cipher mode of operation
	// Using the cipher feedback (CFB) mode here.
	// CBCEncrypter also available.
	cfb := cipher.NewCFBEncrypter(block, iv)
	// Generate the encrypted message and store it
	// in the remaining bytes after the IV nonce
	cfb.XORKeyStream(cipherText[aes.BlockSize:], message)

	return cipherText, nil
}

// AES decryption
func decrypt(key, cipherText []byte) ([]byte, error) {
	// Initialize block cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Separate the IV nonce from the encrypted message bytes
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	// Decrypt the message using the CFB block mode
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(cipherText, cipherText)

	return cipherText, nil
}

func main() {
	// if generate key flag, just output a key to stdout and exit
	keyFile, file, decryptFlag := checkArgs()

	// Load key from file
	keyFileData, err := ioutil.ReadFile(keyFile)
	if err != nil {
		log.Fatal("Unable to read key file contents.", err)
	}

	// Load file to be encrypted or decrypted
	fileData, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("Unable to read key file contents.", err)
	}

	// Perform encryption unless the decryptFlag was provided
	// Outputs to STDOUT. User can redirect output to file.
	if decryptFlag {
		message, err := decrypt(keyFileData, fileData)
		if err != nil {
			log.Fatal("Error decrypting. ", err)
		}
		fmt.Printf("%s", message)
	} else {
		cipherText, err := encrypt(keyFileData, fileData)
		if err != nil {
			log.Fatal("Error encrypting. ", err)
		}
		fmt.Printf("%s", cipherText)
	}
}
