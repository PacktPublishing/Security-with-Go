package main

import "fmt"

func main() {

	goto customLabel

	// Will never get executed because
	// the goto statement will jump right
	// past this line
	fmt.Println("Hello")

customLabel:
	fmt.Println("World")
}
