package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Shape interface {
	Area() float64
}

func getArea(shape Shape) {
	fmt.Println(shape.Area())
}

func main() {

	r := Rectangle{Width: 7, Height: 8}
	c := Circle{Radius: 5}

	getArea(r)
	getArea(c)
}
