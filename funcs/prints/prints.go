package funcs

import "fmt"

func Prints() {
	fmt.Println("Print Functions:")

	// single line print
	fmt.Print("Hello ")

	// multi line print
	fmt.Println("Go!")

	// print with formatting
	var name string = "John Doe"
	var age uint8 = 25

	fmt.Printf("\nName: %v\nAge: %v", name, age)
}