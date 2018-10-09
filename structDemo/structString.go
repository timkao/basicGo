package main

import (
	"fmt"
	"strconv"
)

type TT struct {
	a int
	b float32
	c string
}

func (tn *TT) String() string {
	// %#v 会给出实例的完整输出，包括它的字段（在程序自动生成 Go 代码时也很有用）
	return strconv.Itoa(tn.a) + "/" + fmt.Sprintf("%f", tn.b) + "/" + fmt.Sprintf("%#v", tn.c)
}

func main() {
	fmt.Println("Hello, playground")

	example := &TT{7, -2.35, "abc\tdef"}
	fmt.Printf("%v\n", example)
}
