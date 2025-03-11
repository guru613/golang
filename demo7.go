package main

import "fmt"

type Person struct {
	Name  string
	Email string
	Phone int64
	Age   int
}

func main() {

	p1 := Person{"guru", "@gmail.com", 683637873, 8}

	fmt.Println(p1.Name)
	fmt.Println(p1.Email)
}
