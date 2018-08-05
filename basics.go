package main

import "fmt"

func main() { // entry point of all the codes
	fmt.Printf("Hello, World\n")
	fmt.Printf("Καλημέρα κόσμε; or こんにちは 世界\n")
	myAge1()
	myAge2()
	ageAndWeight()
	nameAgeWeight()
	nameAge()
	multipleNums()
	minOfTwo()
	fmt.Println("variable from other file", beSeen)
}

// go run basics.go variableBasics.go multipleImports.go
