package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
	Name() string
}


type Rectangle struct {
	Width  float64
	Height float64
}

// Rectangle is a Shape cause it implement all the methods of Shape
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Name() string {
	return "Rectangle"
}


// Circle is a Shape cause it implement all the methods of Shape
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Name() string {
	return "Circle"
}


// Triangle is a Shape cause it implement all the methods of Shape
type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}


func (t Triangle) Area() float64 {
	s := (t.SideA + t.SideB + t.SideC) / 2
	return math.Sqrt(s * (s - t.SideA) * (s - t.SideB) * (s - t.SideC))
}

func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

func (t Triangle) Name() string {
	return "Triangle"
}

func PrintShapeInfo(s Shape) {
	fmt.Printf("\n%s Properties:\n", s.Name())
	fmt.Printf("  Area: %.2f\n", s.Area())
	fmt.Printf("  Perimeter: %.2f\n", s.Perimeter())

	switch v := s.(type) {
	case Rectangle:
		fmt.Printf("  Dimensions: %.2f x %.2f\n", v.Width, v.Height)
	case Circle:
		fmt.Printf("  Radius: %.2f\n", v.Radius)
	case Triangle:
		fmt.Printf("  Sides: %.2f, %.2f, %.2f\n", v.SideA, v.SideB, v.SideC)
	}
}


func TotalArea(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}
	triangle := Triangle{SideA: 3, SideB: 4, SideC: 5}
	
	PrintShapeInfo(rect)
	PrintShapeInfo(circle)
	PrintShapeInfo(triangle)
	
	shapes := []Shape{rect, circle, triangle}
	
	fmt.Printf("\n=== Collection of Shapes ===\n")
	for i, shape := range shapes {
		fmt.Printf("%d. %s - Area: %.2f\n", i+1, shape.Name(), shape.Area())
	}
	
	fmt.Printf("\nTotal area of all shapes: %.2f\n", TotalArea(shapes))
	

	fmt.Println("\n=== Type Checking ===")
	
	var s Shape = circle

	if c, ok := s.(Circle); ok {
		fmt.Printf("This is a Circle with radius: %.2f\n", c.Radius)
	}
	

	switch shape := s.(type) {
	case Rectangle:
		fmt.Println("It's a Rectangle!")
	case Circle:
		fmt.Printf("It's a Circle with radius %.2f!\n", shape.Radius)
	case Triangle:
		fmt.Println("It's a Triangle!")
	default:
		fmt.Println("Unknown shape type")
	}
}