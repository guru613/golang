package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("read  file")

	data, err := os.ReadFile("./hi.html")
	if err != nil {
		fmt.Println("error  while reading a file", err)
		return
	}

	fmt.Println("file content:", data)
}
