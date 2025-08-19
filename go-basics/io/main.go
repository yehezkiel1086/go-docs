package main

import "fmt"

func main() {
	var name string
	var age uint8

	fmt.Printf("Enter your name: ")
	fmt.Scanf("%s\n", &name)

	fmt.Printf("Enter your age: ")
	fmt.Scanf("%d", &age)

	fmt.Printf("Welcome, %v (%v)\n", name, age);
}
