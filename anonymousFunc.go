package main

import (
	"fmt"
)

func main() {
	func(num int) {
		fmt.Println(num)
	}(10)

	example := func(num int) int {
		return num + 10
	}

	fmt.Println(example(21))

}
