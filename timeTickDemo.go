package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1e8)
	boom := time.After(5e8)
	// for loop ＋ time.Sleep + channel
	for {
		select {
		case <-tick:
			fmt.Println("Tick!")
		case <-boom:
			fmt.Println("Boom!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}
}
