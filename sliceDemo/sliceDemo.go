/*
	1. 切片（slice）是对数组一个连续片段的引用
	2. 长度可变的数组
	3. cap(s) 就是从 s[0] 到数组末尾的数组长度
	4. 0 <= len(s) <= cap(s)
	5. 切片只能向后移动 切片不能被重新分片以获取数组的前一个元素
	6. 注意 绝对不要用指针指向 slice。切片本身已经是一个引用类型，所以它本身就是一个指针!!
*/

package main

import (
	"fmt"
)

func main() {
	arr1 := [7]int{1, 2, 3, 4, 5, 6, 7}
	slice1 := arr1[1:5]
	for _, value := range slice1 {
		fmt.Println(value)
	}

	/* reslice can only go forward */
	slice1 = slice1[0:6] // start from 2, go all the way to 7
	for _, value := range slice1 {
		fmt.Println(value)
	}
}
