package main

import (
	"fmt"
)

func main() {
	fromSlice := []int{1, 2, 3, 4}
	toSlice := make([]int, 10)

	num := copy(toSlice, fromSlice) // copy return a number and modifiy ToSlice
	fmt.Println(num)
	fmt.Println(toSlice)

	extraSlice := []int{5, 6, 7, 8}
	newSlice := append(toSlice, extraSlice...) // append returns a new slice, destructure a slice
	newSlice2 := append(extraSlice, 9, 10, 11)

	fmt.Println(toSlice)
	fmt.Println(newSlice)
	fmt.Println(newSlice2)
}
