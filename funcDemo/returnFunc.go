package main

import (
	"fmt"
)

func adder() func(int) int {
	var begin int
	return func(delta int) int {
		begin += delta
		return begin
	}
}

func main() {
	var rf = adder()     // var 可配任意值
	fmt.Println(rf(1))   // "begin" is shared
	fmt.Println(rf(10))  // "begin" is shared
	fmt.Println(rf(100)) // "begin" is shared
}
