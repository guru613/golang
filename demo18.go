package main

import "fmt"

func main() {
	number := 10
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d =%d\n", number, i, number*i)
	}
}
