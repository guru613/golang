package main

import "fmt"

type Shape interface {
	area()
}

type circle struct {
	radious float32
}

func (C circle) area() float32 {
	return 3.14 * C.radious * C.radious
}

func main() {
	C := circle{33.4}
	fmt.Println(C.area())
}
