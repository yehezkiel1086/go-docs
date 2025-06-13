package main

import "fmt"

// func as param
func subTwice(sub func(param1 uint8, param2 uint8) uint8, a uint8, b uint8) uint8 {
	return sub(sub(a, b), b)
}

// closure
func newCounter(a uint8) func() uint8 {
	inc := func() uint8 {
		a++
		return a
	}
	return inc // returns a function
}

func main() {
	// func as var
	add := func(a uint8, b uint8) uint8 {
		return a + b
	}
	sub := func(a uint8, b uint8) uint8 {
		return a - b
	}

	// immediately invoked function
	multiplied := func(a uint8, b uint8) uint8 {
		return a * b
	}(3, 4)

	fmt.Println(add(1, 2))
	fmt.Println(subTwice(sub, 3, 1))
	fmt.Println(multiplied)
	fmt.Println(newCounter(10)()) // other iife implementation
}
