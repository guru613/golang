package main

import "fmt"

func main() {
	mymap := map[int]string{
		10: "guru",
		20: "raj",
		30: "hh",
		40: "hkejw",
	}

	delete(mymap, 10)
	fmt.Println(mymap)

}
