package main

import (
	"fmt"
)

func main() {
	// var name string
	// fmt.Println("enter the name please")
	// fmt.Scan(&name)

	// fmt.Println("my name is", name)

	var name1 string
	var age int
	fmt.Println("enter the name and age like guru 20")
	fmt.Scanf("%s %d", &name1, &age)
	fmt.Printf("Name:%s,Age: %d\n", name1, age)

}
