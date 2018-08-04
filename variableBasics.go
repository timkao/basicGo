package main

import "fmt"

func myAge1() {
	var age int // variables need to declare type
	fmt.Println("my age is", age)
	fmt.Println("-----------------")
}

func myAge2() {
	var age = 10 // declaration of type is omitted when it is given a value
	fmt.Println("my age is", age)
	age = 3
	fmt.Println("my age is", age)
	age = 46
	fmt.Println("my age is", age)
	fmt.Println("-----------------")
}

func ageAndWeight() {
	var age, weight int = 35, 55
	fmt.Println("my age is", age)
	fmt.Println("my weight is", weight)
	fmt.Println("-----------------")
}
