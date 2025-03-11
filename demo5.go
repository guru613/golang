package main

import "fmt"

func main() {
	var arr = [10]int{1, 3, 4, 2, 5, 76, 7}
	evencount := 0
	oddcount := 0

	for _, val := range arr {
		if val%2 == 0 {
			evencount++
		} else {
			oddcount++
		}

	}
	fmt.Println(evencount)
	fmt.Println(oddcount)

}
