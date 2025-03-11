package main

import "fmt"

type User struct {
	Name   string
	Age    int
	salary float64
}

func (U User) demo() {
	fmt.Println("my name is ", U.Name)

}

func main() {

	U := User{"guru", 12, 234.45}
	U.demo()

}
