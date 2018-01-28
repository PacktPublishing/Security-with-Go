package main

import (
	"fmt"
	"reflect"
)

func main() {
	myInt := 42
	intPointer := &myInt

	fmt.Println(reflect.TypeOf(intPointer))
	fmt.Println(intPointer)
	fmt.Println(*intPointer)
}
