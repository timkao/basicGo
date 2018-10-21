package main

import "fmt"

/*
	对于同一个通道，发送操作（协程或者函数中的），在接收者准备好之前是阻塞的
*/
func main() {
	ch1 := make(chan int)

	go pump(ch1)
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)

}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}
