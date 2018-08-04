package main

import (
	"fmt"
	"math"
)

func minOfTwo() {
	a, b := 145.8, 543.8
	c := math.Min(a, b)
	fmt.Println("the min number is", c)
	fmt.Println("-----------------")
}
