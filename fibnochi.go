package main

import (
	"fmt"
)

func fib(num int) (result int, index int) {
	if num == 0 || num == 1 {
		return 1, num
	} else {
		result1, _ := fib(num - 1)
		result2, _ := fib(num - 2)
		return result1 + result2, num
	}
}

func main() {
	result, index := fib(4)

	fmt.Println(result, index)
}
