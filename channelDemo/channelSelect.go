package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go pump(ch1)
	go pump(ch2)
	go suck(ch1, ch2)
	time.Sleep(1e9)

}

func pump(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func suck(ch1, ch2 chan int) {

	for {

		select {
		case v := <-ch1:
			fmt.Printf("Receive from channel 1: %v\n", v)
		case v := <-ch2:
			fmt.Printf("Receive from channel 2: %v\n", v)
		}
	}
}
