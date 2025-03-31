package main

import "fmt"

func main() {
	var (
		name string
		age uint8
		married bool
	)

	const (
    Pi float64 = 3.14
    E float64 = 2.718
	)


	name, age, married = "Ben", 23, false
	hobby, job := "Security", "Software Dev"

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(married)
	fmt.Println(hobby, job)

	fmt.Println("Pi and E:", Pi, E)
}