package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
	https://stackoverflow.com/questions/43887759/read-file-from-relative-path-with-different-callers
	The relative path is in fact not dependent on where the source file is located
	(such as Main.go or CleanupUtil_test.go), but where you execute it from
*/
func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, stringErr := inputReader.ReadString('\n')
		fmt.Printf(inputString)
		if stringErr == io.EOF {
			return
		}
	}

}
