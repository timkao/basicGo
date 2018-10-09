package main

import (
	"fmt"
)

func main() {
	slice1 := make([]int, 10, 50) // "make" returns a value with the type
	slice2 := new([50]int)[0:10]  // "new" return a pointer pointing to an instance of the type with the initial value

	fmt.Println(cap(slice1))
	fmt.Println(cap(slice2))

}
