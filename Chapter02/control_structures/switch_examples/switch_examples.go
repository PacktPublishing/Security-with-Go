package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := 42

	switch x {
	case 25:
		fmt.Println("X is 25")
	case 42:
		fmt.Println("X is the magical 42")
		// Fallthrough will continue to next case
		fallthrough
	case 100:
		fmt.Println("X is 100")
	case 1000:
		fmt.Println("X is 1000")
	default:
		fmt.Println("X is something else.")
	}

	// Like the if statement a statement
	// can be put in front of the switched variable
	switch r := rand.Int(); r {
	case r % 2:
		fmt.Println("Random number r is even.")
	default:
		fmt.Println("Random number r is odd.")
	}
	// r is no longer available after the switch statement
}
