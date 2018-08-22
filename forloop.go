package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 6; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("G")
		}
		fmt.Println("")
	}
	fmt.Println("---------")

	for i := 0; i < 15; i++ {
		fmt.Println(i)
	}
	fmt.Println("---------")

	// 可以在循环中同时使用多个计数器
	for i, j := 0, 20; i < j; i, j = i+1, j-1 {
		fmt.Println(i, j)
	}

	var num = 10

	for num < 20 {
		fmt.Println(num)
		num++
	}

}
