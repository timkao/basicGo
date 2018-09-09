/*
	1. key existence
	2. pair deletion
	3. iterate through map
*/

package main

import (
	"fmt"
)

func main() {
	demoMap := map[string]int{"wholeFood": 5, "sevenEleven": 4, "walmart": 4}
	demoMap["kmart"] = 3
	demoMap["hmart"] = 3

	// map 默认是无序的，不管是按照 key 还是按照 value 默认都不排序
	for key := range demoMap {
		fmt.Printf("%s is here\n", key)
	}

	if _, exists := demoMap["walmart"]; exists {
		fmt.Println("found walmart, perpare to delete it")
		delete(demoMap, "walmart")
		fmt.Println("deleted it")
	}

	for key, value := range demoMap {
		fmt.Printf("%s rating is %d\n", key, value)
	}

}
