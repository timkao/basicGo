package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var inputReader *bufio.Reader
	fmt.Println("Please enter a password: ")
	inputReader = bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Println("Some thing is wrong!")
	}
	switch input {
	case "123\n":
		fmt.Println("too eady")
	case "abc\n":
		fmt.Println("what is this")
	default:
		fmt.Println("cannot recognize it")
	}

}
