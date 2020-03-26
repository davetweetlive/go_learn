package main

import (
	"fmt"
	"go_learn/geometry/circle"
	"go_learn/geometry/rectangle"
)

func main() {
	// Implementation of packages
	fmt.Println("Circle")
	fmt.Println("Area: ", circle.Area(2))
	fmt.Println("Perimeter: ", circle.Perimeter(2))
	fmt.Println("Rectangle")
	fmt.Println("Area: ", rectangle.Area(2, 5))
	fmt.Println("Perimeter: ", rectangle.Perimeter(2, 5))
}
