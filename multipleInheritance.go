package main

import (
	"fmt"
)

type pen struct {
	name string
}

func (penIn *pen) write() {
	fmt.Println("Pen can write")
}

type recorder struct {
	price float32
}

func (recoIn *recorder) record() {
	fmt.Println("recorder does record")
}

type recordPen struct {
	pen
	recorder
}

func main() {
	fmt.Println("Hello, playground")

	example := new(recordPen)
	example.write()
	example.record()

}
