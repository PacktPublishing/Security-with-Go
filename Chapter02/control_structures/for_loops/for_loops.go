package main

import (
	"fmt"
)

func main() {
	// Basic for loop
	for i := 0; i < 3; i++ {
		fmt.Println("i:", i)
	}

	// For used as a while loop
	n := 5
	for n < 10 {
		fmt.Println(n)
		n++
	}
}
