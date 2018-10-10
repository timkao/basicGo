package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	outputFile, outputFileError := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outputFileError != nil {
		fmt.Println("An error occurred during file opening")
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"

	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}
