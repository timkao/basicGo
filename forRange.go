package main

import (
	"fmt"
)

/*
	这是 Go 特有的一种的迭代结构，您会发现它在许多情况下都非常有用。
	它可以迭代任何一个集合（包括数组和 map，详见第 7 和 8 章）。
	语法上很类似其它语言中 foreach 语句，但您依旧可以获得每次迭代所对应的索引。
	一般形式为：for ix, val := range collection { }

	ix is pointer
	while val is a "copy" of value
*/

func main() {
	var str string = "abcdefghijklmn"

	for pos, char := range str {
		fmt.Printf("the %d char is %c\n", pos, char)
	}

	s := ""
	for s != "aaaaa" {
		fmt.Println("Value of s:", s)
		s = s + "a"
	}

	/* 如果你只需要索引，你可以忽略第二个变量 */
	seasons := [...]string{"spring", "summer", "fall", "winter"}
	for idx := range seasons {
		fmt.Println(idx)
	}

}
