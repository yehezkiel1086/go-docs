package calculator_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"ginkgo-calculator/calculator"
)

var _ = Describe("Calculator", func() {
	var calc *calculator.Calculator

	BeforeEach(func() {
		calc = calculator.NewCalculator(8)
	})

	Describe("Add", func() {
		It("Should add the base number with added number", func() {
			// act
			calc.Add(2)

			// assertion
			Expect(calc.Result()).To(Equal(float64(10)))
		})
	})

	Describe("Subtract", func() {
		It("Should subtract the base number with subtracted number", func() {
			calc.Subtract(2)

			Expect(calc.Result()).To(Equal(float64(6)))
		})
	})

	Describe("Multiply", func() {
		It("Should multiply the base number with multiplied number", func() {
			calc.Multiply(3)

			Expect(calc.Result()).To(Equal(float64(24)))
		})
	})

	Describe("Divide", func() {
		It("Should divide the base number with divided number", func() {
			calc.Divide(2)

			Expect(calc.Result()).To(Equal(float64(4)))
		})
	})

	AfterEach(func() {
		calc.Reset()
	})
})
