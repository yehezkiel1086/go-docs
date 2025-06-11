package main

import "fmt"

func main() {
	confName := "Go Conference" // syntactic sugar
	const confTickets uint8 = 50
	var remTickets uint8 = 50

	fmt.Printf("VAR_TYPES: confTickets is %T, remTickets is %T, confName is %T\n", confTickets, remTickets, confName)
	
	fmt.Println("Welcome to", confName, "🥳")
	fmt.Println("Note: Get your tickets here to attend.")
	fmt.Printf("There are %v left out of %v\n", remTickets, confTickets)

	var userName string
	var userTickets uint8
	var userEmail string

	fmt.Printf("\nEnter your first name: ")
	fmt.Scan(&userName)
	fmt.Printf("Enter your email: ")
	fmt.Scan(&userEmail)
	fmt.Printf("Enter total tickets to order: ")
	fmt.Scan(&userTickets)

	fmt.Printf("\nThanks, %v for ordering %v tickets!\n", userName, userTickets)
	fmt.Printf("You'll receive ticket at your email: %v\n\n", userEmail)
	remTickets -= userTickets
	fmt.Printf("Remaining tickets: %v\n", remTickets)
}
