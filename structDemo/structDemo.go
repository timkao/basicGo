package main

import (
	"fmt"
)

/*
	1. create a struct
	2. new method returns a "pointer"
*/

/*
	Go 语言中，结构体和它所包含的数据在内存中是以连续块的形式存在的，即使结构体中嵌套有其他的结构体，
	这在性能上带来了很大的优势。不像 Java 中的引用类型，一个对象和它里面包含的对象可能会在不同的内存空间中，
*/

func main() {
	var T struct{ a, b int }

	example := T
	example.a = 100
	example.b = 20

	fmt.Printf("A is %d. B is %d\n", example.a, example.b)
	fmt.Println(example)

	type structEx struct {
		num      int
		fraction float32
		str      string
	}

	ex := new(structEx) // to use "new". the sturct needs to be initialized with "type"
	ex.num = 20
	ex.fraction = 11.11
	ex.str = "This is an example"

	// ex is a pointer in this case
	fmt.Println(ex.num)
	fmt.Println(ex.fraction)
	fmt.Println(ex.str)
	fmt.Println(*ex)
	fmt.Println(ex)

	// ex2 is a value
	ex2 := structEx{40, 22.22, "another way to init"}
	fmt.Println(ex2.num)
	fmt.Println(ex2.fraction)
	fmt.Println(ex2.str)
	fmt.Println(ex2)

	/*
		Go 中的类型转换遵循严格的规则。当为结构体定义了一个 alias 类型时，
		此结构体类型和它的 alias 类型都有相同的底层类型，它们可以如示例 10.3 那样互相转换，
		同时需要注意其中非法赋值或转换引起的编译错误
	*/

	type num1 struct {
		num float32
	}

	type num2 num1 // alias type
	a := num1{5.0}
	b := num2{10.0}
	var c = num1(b) // convert num2 to num1
	fmt.Println(a, b, c)

}
