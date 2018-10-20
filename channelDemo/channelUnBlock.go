package main

import (
	"fmt"
	"time"
)

/*
	若使用通道的缓冲，你的程序会在“请求”激增的时候表现更好：更具弹性，专业术语叫：更具有伸缩性（scalable。
	要在首要位置使用无缓冲通道来设计算法，只在不确定的情况下使用缓冲。
*/

func main() {
	ch1 := make(chan int)
	go pump(ch1)
	go suck(ch1)

	time.Sleep(1e9)

}

func pump(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
