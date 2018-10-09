package main

import (
	"fmt"
)

type Exfunc struct {
	num1 int
	num2 int
}

func (ex *Exfunc) Add(num3 int) int {
	return ex.num1 + ex.num2 + num3
}

func main() {
	example := new(Exfunc)
	example.num1 = 1
	example.num2 = 2

	fmt.Println(example.Add(3))

}
