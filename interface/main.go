package main

import (
	"fmt"
	"math"
)

// Shape is an interface that defines the methods any shape should have
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle is a struct representing a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle is a struct representing a circle
type Circle struct {
	Radius float64
}

// Implementing the methods of the Shape interface for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

// Implementing the methods of the Shape interface for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func main() {
	rect := Rectangle{Width: 6, Height: 3}
	circ := Circle{Radius: 3.5}

	fmt.Println("Rectangle:")
	printShapeInfo(rect)

	fmt.Println("\nCircle:")
	printShapeInfo(circ)
}
