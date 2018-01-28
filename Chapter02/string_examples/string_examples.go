package main

import "fmt"

func main() {
	// Long form assignment
	var myText = "test string 1"

	// Short form assignment
	myText2 := "test string 2"

	// Multiline string
	myText3 := `long string
spanning multiple
lines`

	fmt.Println(myText)
	fmt.Println(myText2)
	fmt.Println(myText3)
}
