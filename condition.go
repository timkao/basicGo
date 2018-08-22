package main

import (
	"fmt"
	"runtime"
)

func main() {
	if runtime.GOOS == "window" {
		fmt.Println("This is window system")
	} else {
		fmt.Printf("this is %s \n", runtime.GOOS)
	}

	// if 可以包含一个初始化语句（如：给一个变量赋值）。这种写法具有固定的格式（在初始化语句后方必须加上分号)
	// 使用简短方式 := 声明的变量的作用域只存在于 if 结构中
	max := 10
	if val := 11; val > max {
		fmt.Println("Val is initialized in the if clause")
	}

}
