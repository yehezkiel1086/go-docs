package main

import (
	"fmt"
	"ginkgo-calculator/calculator"
)

func main() {
	var calc *calculator.Calculator

	calc = calculator.NewCalculator(3)

	calc.Add(2)

	fmt.Println(calc.Result())

	calc.Reset()
}
