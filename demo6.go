package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4}
	arr2 := []int{5, 6, 7, 8}
	arr = append(arr, arr2...)
	fmt.Println(arr)
}
