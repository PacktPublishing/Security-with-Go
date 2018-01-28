package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func printUsage() {
	fmt.Println("Usage: " + os.Args[0] + " <password>")
	fmt.Println("Example: " + os.Args[0] + " Password1!")
}
func checkArgs() string {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
	return os.Args[1]
}

// secretKey should be unique, protected, private,
// and not hard-coded like this. Store in environment var
// or in a secure configuration file.
// This is an arbitrary key that should only be used for example purposes.
var secretKey = "neictr98y85klfgneghre"

// Create a salt string with 32 bytes of crypto/rand data
func generateSalt() string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(randomBytes)
}

// Hash a password with the salt
func hashPassword(plainText string, salt string) string {
	hash := hmac.New(sha256.New, []byte(secretKey))
	io.WriteString(hash, plainText+salt)
	hashedValue := hash.Sum(nil)
	return hex.EncodeToString(hashedValue)
}

func main() {
	// Get the password from command line argument
	password := checkArgs()
	salt := generateSalt()
	hashedPassword := hashPassword(password, salt)
	fmt.Println("Password: " + password)
	fmt.Println("Salt: " + salt)
	fmt.Println("Hashed password: " + hashedPassword)
}
