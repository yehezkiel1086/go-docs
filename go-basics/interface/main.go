package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	Area() float64
	Circumference() float64
}

type ThreeDimension interface {
	Geometry
	Volume() float64
}

type Cube struct {
    Side float64
}

func (c Cube) Volume() float64 {
    return math.Pow(c.Side, 3)
}

func (c Cube) Area() float64 {
    return math.Pow(c.Side, 2) * 6
}

func (c Cube) Circumference() float64 {
    return c.Side * 12
}

type Rectangle struct {
	Side float64
}

func (s Rectangle) Area() float64 {
	return s.Side * s.Side
}

func (s Rectangle) Circumference() float64 {
	return 4 * s.Side
}

type Circle struct {
	Radius float64
}

func (s Circle) Area() float64 {
	return s.Radius * s.Radius * math.Pi
}

func (s Circle) Circumference() float64 {
	return 2 * math.Pi * s.Radius
}

func main() {
	var s Geometry = &Rectangle{
		Side: 4,
	}
	var c Geometry = &Circle{
		Radius: 7,
	}
	var cube ThreeDimension = &Cube{
		Side: 10,
	}

	fmt.Println("Square:")
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Circumference: %.2f\n\n", s.Circumference())

	fmt.Println("Circle:")
	fmt.Printf("Area: %.2f\n", c.Area())
	fmt.Printf("Circumference: %.2f\n\n", c.Circumference())

	fmt.Println("Cube:")
	fmt.Printf("Area: %.2f\n", cube.Area())
	fmt.Printf("Circumference: %.2f\n", cube.Circumference())
	fmt.Printf("Volume: %.2f\n", cube.Volume())
}
