package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Nil maps will cause runtime panic if used
	// without being initialized with make()
	var intToStringMap map[int]string
	var stringToIntMap map[string]int
	fmt.Println(reflect.TypeOf(intToStringMap))
	fmt.Println(reflect.TypeOf(stringToIntMap))

	// Initialize a map using make
	map1 := make(map[string]string)
	map1["Key Example"] = "Value Example"
	map1["Red"] = "FF0000"
	fmt.Println(map1)

	// Initialize a map with literal values
	map2 := map[int]bool{
		4:  false,
		6:  false,
		42: true,
	}

	// Access individual elements using the key
	fmt.Println(map1["Red"])
	fmt.Println(map2[42])

	// Use range to iterate through maps
	for key, value := range map2 {
		fmt.Printf("%d: %t\n", key, value)
	}

}
