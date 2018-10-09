package main

import (
	"log"
	"runtime"
)

func main() {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line) // useful for debug
	}
	where()
	where()
	log.SetFlags(log.Llongfile)
	log.Print("This is to check flag functionality")

	var where2 = log.Print
	where2()
	where2()
}
