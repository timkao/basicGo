package main

import "fmt"

type printInterface interface {
	print()
}

type myInterface interface {
	print()
	start()
}

type exampleStruct struct {
	num int
}

func (p *exampleStruct) print() {
	fmt.Println("Hello World!")
}

func (p *exampleStruct) start() {
	fmt.Println("Start!")
}

func iToI(from myInterface) {
	from.(printInterface).print()
}

func main() {
	example := &exampleStruct{10}

	iToI(example)
}
