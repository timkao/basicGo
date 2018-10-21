package main

import (
	"fmt"
)

var resume chan int

func createIntGenerator() chan int {
	yield := make(chan int)
	count := 0
	go func() {
		for {
			yield <- count
			count++
		}
	}()

	return yield
}

func generateInt() int {
	return <-resume
}

func main() {
	resume = createIntGenerator()
	//利用 block 的特性來設計 generator
	fmt.Println(generateInt())
	fmt.Println(generateInt())
	fmt.Println(generateInt())
}
