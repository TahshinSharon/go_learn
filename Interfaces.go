package main

import (
	"fmt"
	"math"
)

// Define an interface for geometric shapes
type geometry interface {
	area() float64  // Method to calculate area
	perim() float64 // Method to calculate perimeter
}

// Define a struct for rectangles
type rect struct {
	width, height float64
}

// Define a struct for circles
type circle struct {
	radius float64
}

// Implement the geometry interface for rect
func (r rect) area() float64 {
	return r.width * r.height // Area of rectangle = width * height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height // Perimeter of rectangle = 2 * (width + height)
}

// Implement the geometry interface for circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius // Area of circle = π * radius²
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius // Perimeter of circle = 2 * π * radius
}

// Function that takes any geometry and prints its details
func measure(g geometry) {
	fmt.Println(g)         // Print the geometry's details
	fmt.Println(g.area())  // Print the area
	fmt.Println(g.perim()) // Print the perimeter
}

func main() {
	// Create instances of rect and circle
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// Use the measure function for both shapes
	measure(r)
	measure(c)
}
