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
	var age, weight int = 35, 55 // declare multiple variables at the same time
	fmt.Println("my age is", age)
	fmt.Println("my weight is", weight)
	fmt.Println("-----------------")
}

func nameAgeWeight() { // declare multiple variables with different types at the same time
	var (
		name   = "Tim"
		age    = 35
		weight = 55
	)
	fmt.Println("my name is", name)
	fmt.Println("my age is", age)
	fmt.Println("my weight is", weight)
	fmt.Println("-----------------")
}

func nameAge() { // short hand declaration
	name, age := "Tim", 35
	fmt.Println("This is my name", name)
	fmt.Println("This is my age", age)
	fmt.Println("-----------------")
}

/*
	Short hand syntax can only be used when at least one of the variables
	in the left side of := is newly declared. Consider the following program
*/
func multipleNums() {
	a, b := 20, 40
	fmt.Println("a is", a, "b is", b)
	b, c := 50, 60
	fmt.Println("b is", b, "c is", c)
	b, c = 80, 90
	fmt.Println("b is", b, "c is", c)
	fmt.Println("-----------------")
}

var beSeen = 10000
