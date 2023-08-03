package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Square struct {
	SideLength float64
}

func (s Square) Area() float64 {
	return s.SideLength * s.SideLength
}

func PrintArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	rectangle := Rectangle{Width: 5, Height: 3}
	square := Square{SideLength: 4}

	PrintArea(rectangle) // Output: Area: 15
	PrintArea(square)    // Output: Area: 16
}

// type Rectangle struct {
//     Width  float64
//     Height float64
// }

// func (r Rectangle) Area() float64 {
//     return r.Width * r.Height
// }

// type Square struct {
//     SideLength float64
// }

// func (s Square) Area() float64 {
//     return s.SideLength * s.SideLength
// }

// func PrintArea(rectangle Rectangle) {
//     fmt.Println("Area:", rectangle.Area())
// }

// func main() {
//     rectangle := Rectangle{Width: 5, Height: 3}
//     square := Square{SideLength: 4}

//     PrintArea(rectangle) // Output: Area: 15
//     PrintArea(square)   // Incorrect output: Area: 16 (should be 16, but it's not)
// }
