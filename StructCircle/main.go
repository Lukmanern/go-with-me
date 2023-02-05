package main

import (
	"fmt"
	"math"
)

// Define a circle type
// that has a radius field.
type Circle struct {
	Radius float64
}

// Define a method on the Circle type that calculates
// and returns the area of the circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main() {
	// Create a new Circle with a radius of 5.
	c := Circle{Radius: 30}

	// Print the area of the circle.
	fmt.Printf("Area of circle with radius %v: %v\n", c.Radius, c.Area())
}