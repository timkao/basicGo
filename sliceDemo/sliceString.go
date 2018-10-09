package main

import (
	"fmt"
)

func main() {
	s := "This is a test"
	c1 := []byte(s)
	fmt.Println(c1)

	c2 := []int32(s) // for Unicode
	fmt.Println(c2)

	c3 := []rune(s) // when there are non-english characters
	fmt.Println(c3)

	s2 := "Should be in the back"
	c4 := append(c1, s2...)
	fmt.Println(c4)

}
