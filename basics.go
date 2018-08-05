package main

import "fmt"

func main() { // entry point of all the codes
	fmt.Printf("Hello, World\n")
	// other languages work perfectly
	fmt.Printf("Καλημέρα κόσμε; or こんにちは 世界\n")

	// demonstrate ways to set variables
	myAge1()
	myAge2()
	ageAndWeight()
	nameAgeWeight()
	nameAge()
	multipleNums()
	fmt.Println("variable from other file", beSeen)

	// multiple import format
	minOfTwo()

	// demonstrate differnet data types
	fmt.Println("----- Types in Go -------")
	demonstrateBool()
	intDemonstration()
	floatDemonstration()
	stringDemonstration()
	convertDemonstration()

	// demonstrate constant
	fmt.Println("----- Constant in Go -------")
	stringConst()
	strongTypingPolicy()
	numberConst()
}

// go run basics.go variableBasics.go multipleImports.go
