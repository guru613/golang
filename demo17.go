package main

import "fmt"

func main() {

	var num int
	fmt.Println("enter the number:")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Println("its a even number")
	} else {
		fmt.Println("its not a even number")
	}
}
