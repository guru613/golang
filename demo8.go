package main

import "fmt"

func main() {

	// var mymap = map[string]string{

	// 	"guru":  "guru1",
	// 	"hi":    "bye",
	// 	"apple": "banana",
	// }
	// fmt.Println(mymap)

	mymap := make(map[int]string)

	mymap[10] = "apple"
	mymap[20] = "banana"
	mymap[30] = "orange"

	fmt.Println(mymap)

	for _, val := range mymap {
		if val == "apple" {
			fmt.Println(val)
		}
	}

	value, exists := mymap[10]

	if exists {
		fmt.Println(value)
	} else {
		fmt.Println("value not found")
	}

}
