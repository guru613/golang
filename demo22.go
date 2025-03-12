package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("./guru.html")

	if err != nil {
		
		fmt.Println("error while creating file", err)
		return
	}
	defer file.Close()

	//writing for file
	_, err = file.WriteString("hi guru")

	if err != nil {
		fmt.Println("error while writing file", err)
		return
	}

	fmt.Println("file written successfully")

}
