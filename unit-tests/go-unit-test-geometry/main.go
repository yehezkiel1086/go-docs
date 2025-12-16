package main

import (
	"fmt"
	"math"
)

type Cube struct {
	Side float64
}

func (c *Cube) Area() float64 {
	return math.Pow(c.Side, 2) * 6
}

func (c *Cube) Circumference() float64 {
	return c.Side * 12
}

func (c *Cube) Volume() float64 {
	return math.Pow(c.Side, 3)
}

func main() {
	cube := Cube{
		Side: 2,
	}

	area := cube.Area()
	vol := cube.Volume()
	circ := cube.Circumference()

	fmt.Println(area, vol, circ)
}