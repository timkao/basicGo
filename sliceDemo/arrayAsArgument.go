/*
	把一个大数组传递给函数会消耗很多内存。有两种方法可以避免这种现象：
		1. 传递数组的指针
		2. 使用数组的切片
*/

package main

import (
	"fmt"
)

func main() {
	var arr1 = [...]float64{2.2, 4.1, 2.5}
	fmt.Println(sumFloat(&arr1)) // pass pointer
}

func sumFloat(numP *[3]float64) float64 {
	result := 0.0
	for _, value := range numP {
		result += value
	}
	return result
}
