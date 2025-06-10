package main

import (
	"fmt"
	"math"
)

// geometry defines the behaviors required for a shape
// Any type implementing these methods satisfies this interface
// Example: Both rectangle and circle implement this interface
type geometry interface {
	area() float64      // Calculate shape's area
	perimeter() float64 // Calculate shape's perimeter
}

// rectangle represents a 4-sided shape with length and width
type rectangle struct {
	length, width float64
}

// circle represents a round shape defined by its radius
type circle struct {
	radius float64
}

// Implementation of geometry interface for rectangle

// area calculates rectangle area (length × width)
// Example: rectangle{5, 3}.area() returns 15
func (r rectangle) area() float64 {
	return r.length * r.width
}

// perimeter calculates rectangle perimeter (2 × (length + width))
// Example: rectangle{5, 3}.perimeter() returns 16
func (r rectangle) perimeter() float64 {
	return 2 * (r.length + r.width)
}

// Implementation of geometry interface for circle

// area calculates circle area (π × radius²)
// Example: circle{4}.area() returns ~50.27
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// perimeter calculates circle circumference (2 × π × radius)
// Example: circle{4}.perimeter() returns ~25.13
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	// Create shape instances
	r := rectangle{length: 5, width: 3}
	c := circle{radius: 4}

	// Store different shapes in the same slice through interface
	// Example: Both rectangle and circle can be treated as geometry
	shapes := []geometry{r, c}

	// Polymorphism in action - each shape uses its own implementation
	// Example: shape.area() calls rectangle.area() or circle.area() as appropriate
	for _, shape := range shapes {
		fmt.Printf("Shape: %T, Area: %.2f, Perimeter: %.2f\n",
			shape, shape.area(), shape.perimeter())
	}
}
