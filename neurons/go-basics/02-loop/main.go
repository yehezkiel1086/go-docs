package main

import "fmt"

func main() {
	// if with short statement
	if i := 0; i > 5 {
		fmt.Println("More than 5")
	} else {
		fmt.Println("Less than or equals 5")
	}

	// for loop
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("Index: %v\n", i)
	// }

	// conditional loop
	// i := 0
	// for i < 5 {
	// 	fmt.Printf("Index: %v\n", i)
	// 	i++
	// }

	// infinite loop
	for {
		var name string
		fmt.Print("Enter your name: ")
		fmt.Scan(&name)

		fmt.Println("Welcome,", name)
	}
}