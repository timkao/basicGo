package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	pa := &Address{"home", "Taipei", "Taiwan"}
	wa := &Address{"work", "New York", "USA"}
	vc := &VCard{"Tim", "Kao", []*Address{pa, wa}, "flexibility"}

	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s", js)
	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()

	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		fmt.Println("cannot write in file vcard.json")
	}

}
