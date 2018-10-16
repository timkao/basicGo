package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

/*
	the variable inside struct can start from upper case or lower case
	but only the one which is upper case can be exported!
*/

type P struct {
	IntA, IntB, IntC int
	Name             string
}

type Q struct {
	IntA, IntB *int32
	Name       string
}

func main() {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)

	// send the value
	err := enc.Encode(P{3, 4, 5, "Tim"})
	if err != nil {
		log.Fatal("encode error:", err)
	}

	// receive the value
	var result Q
	err = dec.Decode(&result)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Printf("%s: %d, %d\n", result.Name, *result.IntA, *result.IntB)

}
