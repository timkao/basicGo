package main

import (
	"fmt"
)

type Valuable interface {
	getValue() float32
}

type StockOption struct {
	price   float32
	amount  float32
	company string
}

type Car struct {
	price float32
	brand string
}

func (s StockOption) getValue() float32 {
	return s.price * s.amount
}

func (c Car) getValue() float32 {
	return c.price
}

func main() {
	fmt.Println("Hello, playground")

	var example Valuable
	example = StockOption{12.2, 10.0, "Nielsen"}
	fmt.Println(example.getValue())

	example = Car{100.0, "Audi"}
	fmt.Println(example.getValue())
}
