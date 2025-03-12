package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("files in golang")

	file, err := os.Create("./hi.txt")

	if err != nil {
		fmt.Println("error creating with file", err)
		return
	}
	defer file.Close()

	fmt.Println("file created succesfully")

}
