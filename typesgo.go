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

/*
	Type Conversion
	Explicit type conversion is required to assign a variable of one type to another.
*/
func convertDemonstration() {
	i := 57
	j := 66.8
	sum := i + int(j)
	fmt.Println("The sum after conversion is", sum)
}

/*
	Constant
	A string constant like "Hello World" does not have any type. 所以可以任意賦予類
	default type of the string constant is string

	A boolean constatn has the same rule as string constant
	A number constant does not have any type
*/
func stringConst() {
	var name = "Sam"
	fmt.Printf("type %T value %v\n", name, name)
}

func strongTypingPolicy() {
	var defaultName = "Sam"
	type myString string
	var customName myString = "Sam"
	// customName = defaultName   /* this will not be allowed */
	fmt.Printf("Type fo custom name is %T, Type of default name is %T\n", customName, defaultName)
}

func numberConst() {
	const a = 5
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)
}
