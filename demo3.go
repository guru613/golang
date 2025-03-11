package main

import "fmt"

func main() {
	var days = []string{"monday", "sunday", "tuesday"}

	fmt.Println(days)

	fmt.Println(days[0])
	fmt.Println(days[1])
	fmt.Println(days[2])

	days = append(days, "wedbesday")
	fmt.Println(days)
}
