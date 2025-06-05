// Package main demonstrates Go method declarations and usage
package main

import (
	"fmt"
)

// Rectangle represents a four-sided shape with length and width
type Rectangle struct {
	length float64
	width  float64
}

// Shape embeds Rectangle to demonstrate method promotion
type Shape struct {
	Rectangle
}

// MyInt is a custom integer type that demonstrates methods on basic types
type MyInt int

// area calculates and returns the area of a rectangle
// This is a method with a value receiver
func (r Rectangle) area() float64 {
	return r.length * r.width
}

// scale multiplies both dimensions by the given factor
// This is a method with a pointer receiver to modify the original rectangle
func (r *Rectangle) scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

// isPositive checks if the MyInt value is greater than zero
func (i MyInt) isPositive() bool {
	return i > 0
}

func main() {
	// Create a rectangle and use its methods
	rect := Rectangle{length: 5.0, width: 3.0}

	// Access fields and methods of the rectangle
	fmt.Println("Rectangle length:", rect.length)
	fmt.Println("Rectangle width:", rect.width)
	fmt.Println("Rectangle Area:", rect.area())

	// Modify rectangle using a pointer receiver method
	fmt.Println("Scaling rectangle by factor of 2.0")
	rect.scale(2.0)
	fmt.Println("Scaled Rectangle length:", rect.length)
	fmt.Println("Scaled Rectangle width:", rect.width)
	fmt.Println("Scaled Rectangle Area:", rect.area())

	// Demonstrate methods on basic types
	num1 := MyInt(10)
	num2 := MyInt(-5)
	fmt.Println("Is num1 positive?", num1.isPositive())
	fmt.Println("Is num2 positive?", num2.isPositive())

	// Demonstrate method promotion through embedding
	s := Shape{Rectangle: Rectangle{length: 4.0, width: 2.0}}
	fmt.Println("Shape Rectangle Area:", s.area()) // area() is promoted from Rectangle
}
