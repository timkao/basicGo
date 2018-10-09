package main

import (
	"fmt"
)

type innerS struct {
	num int
	string
}

type outerS struct {
	num int
	innerS
	int
}

func main() {
	ex := new(outerS)
	ex.num = 3
	ex.int = 10
	fmt.Printf("outerS has a num %d and int %d\n", ex.num, ex.int)

	ex2 := outerS{6, innerS{12, "test"}, 20}
	fmt.Printf("annoymous variable of innerS is %s", ex2.innerS.string)

}
