package main

import "fmt"

func main() {
	var name, city string

	// Scan
	// fmt.Print("Enter your name and city: ")
	// fmt.Scan(&name, &city)

	// Scanln
	// fmt.Print("Enter your name: ")
	// fmt.Scanln(&name)

	// fmt.Print("Enter your city: ")
	// fmt.Scanln(&city)

	// Scanf
	fmt.Print("Enter your name and city: ")
	fmt.Scanf("%s %s", &name, &city)

	fmt.Printf("Hello, %v! You live in %v\n", name, city)
}