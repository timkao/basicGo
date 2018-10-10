package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := "products.txt"
	outputFile := "copyProducts.txt"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}

	fmt.Printf("%s\n", string(buf))
	err2 := ioutil.WriteFile(outputFile, buf, 0644)
	if err2 != nil {
		panic(err2.Error())
	}

}
