package main

import (
	"fmt"
)

type Any interface{}
type SpecialString string

var exampleTest SpecialString = "Hello"

func main() {
	testFunc := func(example Any) {
		switch t := example.(type) {
		case bool:
			fmt.Printf("Any %v is a boolean type\n", t)
		case int:
			fmt.Printf("Any %v is int type\n", t)
		case string:
			fmt.Printf("Any %v is string type\n", t)
		case SpecialString:
			fmt.Printf("Any %v is special type\n", t)
		default:
			fmt.Println("Unknow Type")
		}
	}
	testFunc(exampleTest)
}
