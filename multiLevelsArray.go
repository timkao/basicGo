/*
  1. demonstrate multi level array
	2. the use of "type"
*/

package main

import (
	"fmt"
)

func main() {

	const (
		WIDTH  = 1920
		HEIGHT = 1080
	)

	type pixel int
	var screen [WIDTH][HEIGHT]pixel // screen := [WIDTH][HEIGHT]pixel is wrong since [WIDTH][HEIGHT]pixel is not a value
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = 0
		}
	}

	fmt.Println(screen)
}
