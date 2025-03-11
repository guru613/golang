package main

import "fmt"

func greeting() {
	fmt.Println("hi good morning")

}

func add(a, b int) int {
	sum := a + b
	return sum
}

func main() {
	greeting()
	fmt.Println(add(1, 2))

}
