package main

import "fmt"

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	arr[2] = 100
	fmt.Println(arr)

	// for i := 1; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }
	// var sum = 0
	// for i := 1; i < len(arr); i++ {
	// 	sum = sum + arr[i]
	// }
	// fmt.Println(sum)

	var max = arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	fmt.Println(max)

}
