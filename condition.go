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
}
