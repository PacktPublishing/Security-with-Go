package main

import (
	"encoding/base64"
	"fmt"
	"github.com/skip2/go-qrcode"
	"log"
	"os"
)

var (
	pngData        []byte
	imageSize      = 256 // Length and width in pixels
	err            error
	outputFilename string
	dataToEncode   string
)

// Check command line arguments. Print usage
// if expected arguments are not present
func checkArgs() {
	if len(os.Args) != 3 {
		fmt.Println(os.Args[0] + `

Generate a QR code. Outputs a PNG file in <outputFilename>.
Also outputs an HTML img tag with the image base64 encoded to STDOUT.

  Usage: ` + os.Args[0] + ` <outputFilename> <data>
  Example: ` + os.Args[0] + ` qrcode.png https://www.devdungeon.com`)
		os.Exit(1)
	}
	// Because these variables were above, at the package level
	// we don't have to return them. The same variables are
	// already accessible in the main() function
	outputFilename = os.Args[1]
	dataToEncode = os.Args[2]
}

func main() {
	checkArgs()

	// Generate raw binary data for PNG
	pngData, err = qrcode.Encode(dataToEncode, qrcode.Medium, imageSize)
	if err != nil {
		log.Fatal("Error generating QR code. ", err)
	}

	// Encode the PNG data with base64 encoding
	encodedPngData := base64.StdEncoding.EncodeToString(pngData)

	// Output base64 encoded image as HTML image tag to STDOUT
	// This img tag can be embedded in an HTML page
	imgTag := "<img src=\"data:image/png;base64," +
		encodedPngData + "\" />"
	fmt.Println(imgTag) // For use in HTML

	// Generate and write to file with one function
	// This is a standalone function. It can be used by itself
	// without any of the above code
	err = qrcode.WriteFile(
		dataToEncode,
		qrcode.Medium,
		imageSize,
		outputFilename,
	)
	if err != nil {
		log.Fatal("Error generating QR code to file. ", err)
	}
}
