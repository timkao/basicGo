package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

/* both pointer and value work */
func (sq Square) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

func main() {
	fmt.Println("Hello, playground")
	var example Shaper       // a variable for the interface
	var squareEx = Square{5} // a value instance
	squareEx.side = 5.0
	example = squareEx // assign the interface to the interface

	if v, ok := example.(Square); ok {
		fmt.Printf("the shaper has %T\n", v)
	}

	var circleEx = new(Circle) // a pointer instance
	circleEx.radius = 3.0
	example = circleEx

	if v, ok := example.(*Circle); ok {
		fmt.Printf("the shaper has %T\n", v)
	}

	switch t := example.(type) {
	case Square:
		fmt.Printf("Type Square %T with value %v\n", t, t)
	case *Circle:
		fmt.Printf("Type Circle %T with value %v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}
}
