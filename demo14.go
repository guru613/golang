package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type Circle struct {
	Radius float32
}
type Rectangule struct {
	height float32
	weidth float32
}

func main() {
	C := Circle{33.4}
	R := Rectangule{22.2, 33.5}

	fmt.Println(C.area())
	fmt.Println(R.area())

}

func (C Circle) area() float32 {
	return math.Pi * C.Radius * C.Radius
}

func (R Rectangule) area() float32 {
	return R.height * R.weidth
}
