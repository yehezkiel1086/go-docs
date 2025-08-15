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

type Rectangle struct {
	Side float64
}

func (r Rectangle) Area() float64 {
	return r.Side * r.Side
}

func (r Rectangle) Circumference() float64 {
	return r.Side * 4
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Circumference() float64 {
	return 2 * 3.14 * c.Radius
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

func main() {
	var shape Geometry
	shape = Rectangle{
		Side: 10,
	}

	fmt.Printf("Area: %v, Circumference: %v\n\n", shape.Area(), shape.Circumference())

	var object ThreeDimension
	object = Cube{
		Side: 10,
	}

	fmt.Println("Luas :", object.Area())
	fmt.Println("Keliling :", object.Circumference())
	fmt.Println("volume : ", object.Volume())
}
