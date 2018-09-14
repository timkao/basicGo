package main

import (
	"fmt"
	"reflect"
)

type TagDemo struct {
	num  int     "How many people?"
	name string  "What is the name of the group?"
	avg  float32 "Average ages"
}

func main() {
	demo := new(TagDemo) // return pointer
	demo.num = 10
	demo.name = "Tim"
	demo.avg = 33.2

	for i := 0; i < 3; i++ {
		refTag(*demo, i)
	}
}

func refTag(demo TagDemo, idx int) { // demo should be value not pointer
	demoType := reflect.TypeOf(demo)
	idxField := demoType.Field(idx)
	fmt.Printf("%v\n", idxField.Tag)
}
