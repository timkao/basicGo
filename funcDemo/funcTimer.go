package main

import (
	"fmt"
	"time"
)

func funcTimer(testFunc func(int) (int, int), num int) time.Duration { // 100% match to the type of fib...
	start := time.Now()
	testFunc(num)
	end := time.Now()
	delta := end.Sub(start)
	return delta
}

func main() {
	fmt.Printf("It tooks %s to run the function\n", funcTimer(fib, 20))
}
