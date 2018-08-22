package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 100

	// common stucture
	switch num {
	case 98, 99:
		fmt.Println("The value is %i", num)
	case 100:
		fmt.Printf("%d is the value\n", num)
	}

	// special case
	switch {
	case num < 90:
		fmt.Println("The number is less than 90")
	case num > 90:
		fmt.Println("The number is bigger than 90")
		fallthrough
	case num == 100:
		fmt.Println("fallThrough works and the number is 100")
	}

	// switch 语句的第三种形式是包含一个初始化语句
	switch result := atoi("121"); {
	case result < 100:
		fmt.Println("smaller than 100")
	case result > 100:
		fmt.Println("bigger than 100")
	}
}

func atoi(s string) (check int) {
	check, _ = strconv.Atoi(s)
	return
}
