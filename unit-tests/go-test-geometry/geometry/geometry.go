package geometry

import "math"

type Cube struct {
	Side float64
}

func (c Cube) Area() float64 {
	return 6 * math.Pow(c.Side, 2)
}

func (c Cube) Circumference() float64 {
	return 12 * c.Side
}

func (c Cube) Volume() float64 {
	return math.Pow(c.Side, 3)
}
