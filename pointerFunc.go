package main

import (
	"fmt"
)

func multiply(a, b int, reply *int) { // reply should be a pointer
	*reply = a * b // the value of reply is modified
}

func main() {
	n := 0
	reply := &n // assign reply to be an address of n
	multiply(10, 5, reply)
	fmt.Println(*reply) // retrieve the value from the pointer
}
