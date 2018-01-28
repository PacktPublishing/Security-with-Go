package main

import "fmt"

func main() {
	var mySlice []int // nil slice

	// Append will
	mySlice = append(mySlice, 1, 2, 3, 4, 5)

	// Appending also works on nil slices.
	// Since nil slices have zero capacity, and have
	// no underlying array, it will create one.
	mySlice = append(mySlice, 1, 2, 3, 4, 5)

	// Individual elements can be accessed from a slice
	// just like an array by using the square bracket operator.
	firstElement := mySlice[0]
	fmt.Println("First element:", firstElement)

	// To get only the second and third element, use:
	subset := mySlice[1:4]
	fmt.Println(subset)

	// To get the full contents of a slice except for the first element, use:
	subset = mySlice[1:]
	fmt.Println(subset)

	// To get the full contents of a slice except for the last element, use:
	subset = mySlice[0 : len(mySlice)-1]
	fmt.Println(subset)

	// To copy a slice, use the copy() function.
	// If you assign one slice to another with the equal operator,
	// the slices will point at the same memory location,
	// and changing one would change both slices.
	slice1 := []int{1, 2, 3, 4}
	slice2 := make([]int, 4)

	// Create a unique copy in memory
	copy(slice2, slice1)

	// Changing one should not affect the other
	slice2[3] = 99
	fmt.Println(slice1)
	fmt.Println(slice2)
}
