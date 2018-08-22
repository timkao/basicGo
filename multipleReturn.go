package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var orig string = "20"

	an, err := strconv.Atoi(orig)

	if err != nil {
		fmt.Printf("orig is not an integer - existing with error\n", orig)
		return
	}

	fmt.Printf("The integer is %d\n", an)

	t, ok := mySqrt(25.0) // without "ok", the system will throw error
	if ok {
		fmt.Println(t)
	}

	n := atoi("10")
	fmt.Println(n)

}

func mySqrt(f float64) (v float64, ok bool) {
	if f < 0 {
		return
	}
	return math.Sqrt(f), true
}

// 相當特別，系統自動知道該返回的值
func atoi(s string) (check int) {
	check, _ = strconv.Atoi(s)
	return
}
