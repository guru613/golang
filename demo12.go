package main

import "fmt"

func main() {
	arr := [12]int{1, 2, 3, 4, 5, 66}

	fmt.Println(arr[0:2])
	fmt.Println(arr[0:1])
	fmt.Println(arr[1:])
}
