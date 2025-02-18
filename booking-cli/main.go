package main

import "fmt"

func main() {
	conferenceName := "Go Conference" // type inference (can't be assigned a type or const)
	const conferenceTickets uint8 = 50
	var remainingTickets uint8 = 50

	// print data types
	fmt.Printf("Data Types:\n- conferenceName is %T\n- remainingTickets is %T\n- conferenceName is %T\n\n", conferenceName, remainingTickets, conferenceTickets)

	// welcome message
	fmt.Printf("Welcome to %v!\n", conferenceName)
	fmt.Printf("Get your tickets here to attend.\n\n")

	// display tickets
	fmt.Printf("Total tickets: %v\nRemaining tickets: %v\n\n", conferenceTickets, remainingTickets)

	// ask user for their name
	var userName string = "John Doe"
	var email string = "johndoe@gmail.com"
	var userTickets int = 2

	fmt.Printf("Enter your name: ")
	fmt.Scan(&userName)

	fmt.Printf("Enter your email: ")
	fmt.Scan(&email)

	fmt.Printf("\n%v booked %v tickets.\n", userName, userTickets)

	fmt.Printf("confirmation email sent at %v.\n\n", email)
}