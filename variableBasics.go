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
	// 这是使用变量的首选形式，但是它只能被用在函数体内，而不可以用于全局变量的声明与赋值。使用操作符 := 可以高效地创建一个新的变量，称之为初始化声明。
	name, age := "Tim", 35
	fmt.Println("This is my name", name)
	fmt.Println("This is my age", age)
	fmt.Println("-----------------")
}

/*
	Short hand syntax can only be used when at least one of the variables
	in the left side of := is newly declared. Consider the following program
*/

/*
这被称为 并行 或 同时 赋值。

如果你想要交换两个变量的值，则可以简单地使用 a, b = b, a。

(在 Go 语言中，这样省去了使用交换函数的必要)

空白标识符 _ 也被用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃。

_ 实际上是一个只写变量，你不能得到它的值。这样做是因为 Go 语言中你必须使用所有被声明的变量，但有时你并不需要使用从一个函数得到的所有返回值。

并行赋值也被用于当一个函数返回多个返回值时，比如这里的 val 和错误 err 是通过调用 Func1 函数同时得到：val, err = Func1(var1)

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

// 全局变量是允许声明但不使用
var BeSeen = 10000
