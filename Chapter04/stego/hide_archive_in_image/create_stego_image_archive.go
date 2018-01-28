package main

import (
	"io"
	"log"
	"os"
)

func main() {

	// Open original file
	firstFile, err := os.Open("test.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer firstFile.Close()

	// Second file
	secondFile, err := os.Open("test.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer secondFile.Close()

	// New file for output
	newFile, err := os.Create("stego_image.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	// Copy the bytes to destination from source
	_, err = io.Copy(newFile, firstFile)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(newFile, secondFile)
	if err != nil {
		log.Fatal(err)
	}

}
