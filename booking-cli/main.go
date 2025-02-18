package main

import "fmt"

func main() {
	conferenceName := "Go Conference" // type inference (can't be assigned a type or const)
	const conferenceTickets uint8 = 50
	var remainingTickets uint8 = 50

	// print data types
	fmt.Printf("Data Types: conferenceName is %T, remainingTickets is %T, conferenceName is %T\n\n", conferenceName, remainingTickets, conferenceTickets)

	// welcome message
	fmt.Printf("Welcome to %v!\n", conferenceName)
	fmt.Printf("Get your tickets here to attend.\n\n")

	// display tickets
	fmt.Printf("Total tickets: %v\nRemaining tickets: %v\n\n", conferenceTickets, remainingTickets)

	// ask user for their name
	var userName string = "Tom"
	var userTickets int = 2
	fmt.Printf("%v booked %v tickets.\n\n", userName, userTickets)
}