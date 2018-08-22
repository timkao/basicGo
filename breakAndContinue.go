package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("-------")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	ex := 0
	for { //since there are no checks, this is an infinite loop
		if ex >= 3 {
			break
		}
		//break out of this for loop when this condition is met
		fmt.Println("Value of ex is:", ex)
		ex++
	}
	fmt.Println("A statement just after for loop.")

	for i := 0; i < 7; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd:", i)
	}

}
