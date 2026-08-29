package main

import "fmt"

type Person struct {
	Name string
	Age  uint
}

type Student struct {
	Grade float64
	Person
}

func main() {
	s := &Student{}
	s.Name = "Yehezkiel"
	s.Age = 24
	s.Grade = 3.6

	fmt.Printf("Name: %s\n", s.Name)
	fmt.Printf("Age: %d\n", s.Age)
	fmt.Printf("Grade: %.2f\n", s.Grade)
}
