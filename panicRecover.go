package main

import (
	"fmt"
)

func badCall() {
	panic("this is a bad call")
}

func test() {
	defer func() {
		if e := recover(); e != nil { // recover returns the error
			fmt.Printf("panicing %s\n", e)
		}
	}()
	badCall()
	fmt.Println("After bad call")
}

func main() {
	fmt.Println("calling test")
	test()
	fmt.Println("complete!")
}
