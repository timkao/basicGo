package main

import (
	"basicsgo/sort"
	"fmt"
)

func intsDemo() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 786, -5467984, 7586}
	sort.SortInts(data) // use convenient method created in the package
	for _, val := range data {
		fmt.Println(val)
	}
}

func stringsDemo() {
	data := []string{"monday", "friday", "tuesday", "wednesday", "sunday", "thursday", "", "saturday"}
	temp := sort.StringArray(data)
	sort.Sort(temp)
	if !sort.IsSorted(temp) {
		panic("fail")
	}
	fmt.Printf("The sorted Array is %v\n", temp)
}

func main() {
	intsDemo()
	stringsDemo()
}
