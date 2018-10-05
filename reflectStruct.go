package main

import (
	"fmt"
	"reflect"
)

type unknown struct {
	str1, str2, str3 string
	num              int
}

func (p unknown) String() string {
	return p.str1 + " - " + p.str2 + " - " + p.str3
}

var secret interface{} = unknown{"Ada", "Go", "Oberon", 10}

func main() {
	value := reflect.ValueOf(secret) // call String()
	typ := reflect.TypeOf(secret)
	fmt.Println(value)
	fmt.Println(typ)
	knd := value.Kind()
	fmt.Println(knd)

	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
	}

	results := value.Method(0).Call(nil)
	fmt.Println(results)

}
