package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "Tim "

	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Good Morning", who)
}
