package main

func fib(num int) (result int, index int) {
	if num == 0 || num == 1 {
		return 1, num
	} else {
		result1, _ := fib(num - 1)
		result2, _ := fib(num - 2)
		return result1 + result2, num
	}
}
