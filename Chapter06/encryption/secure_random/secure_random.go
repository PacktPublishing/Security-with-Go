package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

func main() {
	// Generate a random int
	limit := int64(math.MaxInt64) // Highest random number allowed
	randInt, err := rand.Int(rand.Reader, big.NewInt(limit))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Random int value: ", randInt)

	// Alternatively, you could generate the random bytes
	// and turn them into the specific data type needed.
	// binary.Read() will only read enough bytes to fill the data type
	var number uint32
	err = binary.Read(rand.Reader, binary.BigEndian, &number)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Random uint32 value: ", number)

	// Or just generate a random byte slice
	numBytes := 4
	randomBytes := make([]byte, numBytes)
	rand.Read(randomBytes)
	fmt.Println("Random byte values: ", randomBytes)
}
