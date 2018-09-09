/*
	map 传递给函数的代价很小：在 32 位机器上占 4 个字节，64 位机器上占 8 个字节，无论实际上存储了多少数据。通过 key 在 map 中寻找值是很快的，比线性查找快得多，但是仍然比从数组和切片的索引中直接读取要慢 100 倍；所以如果你很在乎性能的话还是建议用切片来解决问题。
*/

/*
	1. map creation2 - two ways
	2. len with map
	3. map capacity (not limit)
*/

package main

import (
	"fmt"
)

func main() {

	// first way to create map
	priceMap := map[string]float32{"one": 1.11} // create the container
	priceMap["two"] = 4.99
	priceMap["three"] = 100.11
	fmt.Printf("Defined a key value pair directly %f\n", priceMap["one"])
	fmt.Printf("a cup of coffee is %f\n", priceMap["two"])
	fmt.Printf("25 cups of coffee is %f\n", priceMap["three"])

	// second way to create map
	var peopleMap = make(map[int]string)
	peopleMap[23] = "jordan"
	peopleMap[30] = "tim"
	fmt.Printf("We have two members, %s and %s\n", peopleMap[23], peopleMap[30])

	// len
	fmt.Println(len(peopleMap))

	// initialize with a initalize size 10
	languageMap := make(map[string]string, 10)
	languageMap["python"] = "first one"
	fmt.Printf("Python is my %s language\n", languageMap["python"])

}
