package calculator

import "errors"

type Calculator struct {
	base float64
}

func NewCalculator(base float64) *Calculator {
	return &Calculator{
		base,
	}
}

func (c *Calculator) Add(n float64) {
	c.base += n
}

func (c *Calculator) Subtract(n float64) {
	c.base -= n
}

func (c *Calculator) Multiply(n float64) {
	c.base *= n
}

func (c *Calculator) Divide(n float64) error {
	if n == 0 {
		return errors.New("cannot divide by zero")
	}

	c.base /= n
	
	return nil
}

func (c *Calculator) Result() float64 {
	return c.base
}

func (c *Calculator) Reset() {
	c.base = 0
}
