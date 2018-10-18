package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	// 如果 2 个协程需要通信，你必须给他们同一个通道作为参数才行。
	go sendData(ch)
	go getData(ch)

	// main() 等待了 1 秒让两个协程完成，如果不这样，sendData() 就没有机会输出。
	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
}

/*
	會照著進通道的順序來
*/
func getData(ch chan string) {
	var input string
	// 注意：不要使用打印状态来表明通道的发送和接收顺序：由于打印状态和通道实际发生读写的时间延迟会导致和真实发生的顺序不同。

	for {
		input = <-ch
		fmt.Printf("%s, \n", input)
	}
}
