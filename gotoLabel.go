package main

import (
	"fmt"
)

/*
	特别注意
	使用标签和 goto 语句是不被鼓励的：
	它们会很快导致非常糟糕的程序设计，而且总有更加可读的替代方案来实现相同的需求
*/

func main() {

EXAMPLE:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				continue EXAMPLE
			}
			fmt.Println(i, j)
		}
	}

	ex := 0
HERE:
	fmt.Println(ex)
	ex++
	if ex == 5 {
		return
	}
	goto HERE

}
