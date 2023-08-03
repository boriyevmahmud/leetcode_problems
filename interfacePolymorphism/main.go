package main

import "fmt"


type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) area() float64 {
	return r.length * r.width
}

func printArea(s Shape) {
	fmt.Printf("Area of shape is: %.2f\n", s.area())
}

func main() {
	c := Circle{radius: 5.0}
	r := Rectangle{length: 10.0, width: 5.0}

	printArea(c)
	printArea(r)
}
