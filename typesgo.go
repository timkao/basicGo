package main

import (
	"fmt"
	"unsafe"
)

func demonstrateBool() {
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	d := a || b
	fmt.Println("c:", c, "d:", d)
}

/*
	Integer
	"int"
	size will change according to system
	64 bit system will assign 8 bytes int
	32 bit system will assign 4 bytes int

	Unsigned Integer
	"uint"
*/
func intDemonstration() {
	var a int = 89
	b := 95
	fmt.Println("value of a is", a, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))
	fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b))
}

/*
	Float
*/
func floatDemonstration() {
	a, b := 5.67, 8.97
	fmt.Printf("\nType of a is %T, size of a is %d", a, unsafe.Sizeof(a))
	fmt.Printf("\nType of b is %T, size of b is %d\n", b, unsafe.Sizeof(b))
	sum := a + b
	diff := a - b
	fmt.Println("sum", sum, "diff", diff)
}

/*
	String
*/

func stringDemonstration() {
	first := "Tim"
	last := "Kao"
	name := first + " " + last
	fmt.Println("My name is", name)
}
