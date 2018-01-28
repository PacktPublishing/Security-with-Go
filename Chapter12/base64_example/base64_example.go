package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	data := []byte("Test data")

	// Encode bytes to base64 encoded string.
	encodedString := base64.StdEncoding.EncodeToString(data)
	fmt.Printf("%s\n", encodedString)

	// Decode base64 encoded string to bytes.
	decodedData, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		log.Fatal("Error decoding data. ", err)
	}
	fmt.Printf("%s\n", decodedData)
}
