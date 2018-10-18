package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start main function")
	go longWait()  // use "go" for another thread
	go shortWait() // use "go" for another thread
	time.Sleep(10 * 1e9)
	fmt.Println("end of main")
}

func longWait() {
	fmt.Println("inside longWait")
	time.Sleep(5 * 1e9)
	fmt.Println("end of longWait")
}

func shortWait() {
	fmt.Println("inside shortWait")
	time.Sleep(1 * 1e9)
	fmt.Println("end of shortWait")
}
