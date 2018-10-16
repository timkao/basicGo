package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

func main() {
	input := "<person><FirstName>Tim</FirstName><LastName>Kao</LastName></person>"
	inputReader := strings.NewReader(input)
	p := xml.NewDecoder(inputReader)

	for t, err := p.Token(); err == nil; t, err = p.Token() { // one token after another until the end of the string
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			fmt.Printf("Token name: %s\n", name)
		case xml.EndElement:
			fmt.Println("end of token")
		case xml.CharData:
			content := string([]byte(token)) // turn token to []byte then to string
			fmt.Printf("this is the content %s\n", content)
		default:
			fmt.Println("cannt read it")
		}
	}
}
