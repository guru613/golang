package main

import "fmt"

func main() {
	var day string
	fmt.Println("enter the day")
	fmt.Scan(&day)

	switch day {
	case "monday":
		fmt.Println("start of the week")

	case "wednesday":
		
		fmt.Println("middle of the week")

	case "friday":
		fmt.Println("weekend")

	case "sunday":
		fmt.Println(" is  a holiday")

	default:
		fmt.Println("in valid data")
	}

}
