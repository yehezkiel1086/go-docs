package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Circumference() float64
}

type Square struct {
	Side float64
}

func (s *Square) Area() float64 {
	return s.Side * s.Side
}

func (s *Square) Circumference() float64 {
	return s.Side * 4
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (c *Circle) Circumference() float64 {
	return c.Radius * 2 * math.Pi
}

func main() {
	var shape Shape
	shape = &Square{
		Side: 4,
	}

	fmt.Printf("Area: %.2f\n", shape.Area())
	fmt.Printf("Circumference: %.2f\n", shape.Circumference())
}
