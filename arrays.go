package main

import (
	"fmt"
)

/*
	Go 语言中的数组是一种 值类型（不像 C/C++ 中是指向首元素的指针），所以可以通过 new() 来创建： var arr1 = new([5]int)。
	所以在函数中数组作为参数传入时，如 func1(arr2)，会产生一次数组拷贝，func1 方法不会修改原始的数组 arr2。
*/

func main() {
	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}
	/* 如果你想修改原数组，那么 arr2 必须通过&操作符以引用方式传过来，例如 func1(&arr2）*/
	arr2 := [...]int{1, 2, 3}
	fakeArrModifier(arr2)
	fmt.Println(arr2)
	goodArrModifier(&arr2)
	fmt.Println(arr2)
}

func fakeArrModifier(numArr [3]int) {
	numArr[0] = 10000
}

func goodArrModifier(numArr *[3]int) {
	numArr[0] = 100000
}
