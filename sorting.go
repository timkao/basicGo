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

type day struct {
	num  int
	name string
}

type dayArray struct {
	data []*day
}

func (p *dayArray) Len() int {
	return len(p.data)
}

func (p *dayArray) Less(i, j int) bool {
	return p.data[i].num < p.data[j].num
}

func (p *dayArray) Swap(i, j int) {
	p.data[i], p.data[j] = p.data[j], p.data[i]
}

func structDemo() {
	Sunday := day{0, "SUN"}
	Monday := day{1, "MON"}
	Tuesday := day{2, "TUE"}
	Wednesday := day{3, "WED"}
	Thursday := day{4, "THU"}
	Friday := day{5, "FRI"}
	Saturday := day{6, "SAT"}
	data := []*day{&Tuesday, &Thursday, &Wednesday, &Sunday, &Monday, &Friday, &Saturday}
	temp := dayArray{data} // a way to include data in dayArray
	sort.Sort(&temp)       // the receiver of Len here accepts only pointer
	if !sort.IsSorted(&temp) {
		panic("fail")
	}
	for _, d := range temp.data {
		fmt.Printf("%s ", d.name)
	}
}

func main() {
	intsDemo()
	stringsDemo()
	structDemo()
}
