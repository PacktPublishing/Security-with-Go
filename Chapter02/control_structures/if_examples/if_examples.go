package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Int()

	if x < 100 {
		fmt.Println("x is less than 100.")
	}

	if x < 1000 {
		fmt.Println("x is less than 1000.")
	} else if x < 10000 {
		fmt.Println("x is less than 10,000.")
	} else {
		fmt.Println("x is greater than 10,000")
	}

	fmt.Println("x:", x)

	// You can put a statement before the condition
	// The variable scope of n is limited
	if n := rand.Int(); n > 1000 {
		fmt.Println("n is greater than 1000.")
		fmt.Println("n:", n)
	} else {
		fmt.Println("n is not greater than 1000.")
		fmt.Println("n:", n)
	}
	// n is no longer available past the if statement

}
