package main

import (
	"fmt"
)

type rectangle struct {
	long   float32
	height float32
}

func (rect *rectangle) Area() float32 {
	return rect.long * rect.height
}

type pictures struct {
	long   float32
	height float32
	rectangle
}

func main() {
	fmt.Println("Hello, playground")

	rectEx := new(rectangle)
	rectEx.long = 2.5
	rectEx.height = 2.0
	fmt.Println(rectEx.Area())

	pictEx := pictures{3.8, 9.1, rectangle{3.5, 4.0}}
	fmt.Println(pictEx.Area())
}
