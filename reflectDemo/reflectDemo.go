package main

import (
	"fmt"
	"reflect"
)

func main() {
	var example float64 = 3.2
	fmt.Println("type", reflect.TypeOf(example))
	v := reflect.ValueOf(example)
	fmt.Println("value", v)
	fmt.Println("type", v.Kind())
	fmt.Println("value", v.Float())
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)
}
