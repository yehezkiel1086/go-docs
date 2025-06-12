package main

import (
	"fmt"
	"strings"
)

var confName string
const confTickets uint8 = 50
var remTickets uint8 = 50
var bookings []string
// slice
// var bookings []string

func main() {
	confName = "Go Conference" // syntactic sugar

	fmt.Printf("VAR_TYPES: confTickets is %T, remTickets is %T, confName is %T\n", confTickets, remTickets, confName)

	greetUsers()

	// array
	// var bookings [50]string
	// bookings[0] = "Nana"
	// bookings[1] = "Nicole"
	// bookings[2] = "Peter"
	// could also be like this:
	// var bookings = [50]string{"Nana", "Nicole", "Peter"}

	for {
		firstName, lastName, userEmail, userTickets := getUserInput()

		isValidInput := validateInput(firstName, lastName, userEmail, userTickets)

		if isValidInput {
			bookTicket(userTickets, firstName, lastName, userEmail)
			if remTickets == 0 {
				fmt.Println("Ticket is sold out. Come back next year!")
				break
			}
		}
	}
}

func greetUsers() {
	fmt.Println("Welcome to", confName, "🥳")
	fmt.Println("Note: Get your tickets here to attend.")
	fmt.Printf("There are %v left out of %v\n", remTickets, confTickets)
}

func getFirstnames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func validateInput(firstName string, lastName string, userEmail string, userTickets uint8) bool {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(userEmail, "@")
	isValidTicket := userTickets > 0 && userTickets <= remTickets

	if !isValidName {
		fmt.Println("First or last name is too short.")
	}
	if !isValidEmail {
		fmt.Println("Email requires @")
	}
	if !isValidTicket {
		fmt.Println("Invalid tickets order")
	}

	return isValidName && isValidEmail && isValidTicket
}

func getUserInput() (string, string, string, uint8) {
	var firstName string
	var lastName string
	var userTickets uint8
	var userEmail string

	fmt.Printf("\nEnter your first name: ")
	fmt.Scan(&firstName)
	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Printf("Enter your email: ")
	fmt.Scan(&userEmail)
	fmt.Printf("Enter total tickets to order: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, userEmail, userTickets
}

func bookTicket(userTickets uint8, firstName string, lastName string, userEmail string) {
	remTickets -= userTickets
	// bookings[0] = userName // add element to array
	bookings = append(bookings, firstName + " " + lastName) // add element to slice

	// fmt.Printf("\nWhole array: %v\n", bookings)
	// fmt.Printf("Array type: %T\n", bookings)
	// fmt.Printf("Array length: %v\n\n", len(bookings))
	// fmt.Printf("\nWhole slice: %v\n", bookings)
	// fmt.Printf("Slice type: %T\n", bookings)
	// fmt.Printf("Slice length: %v\n\n", len(bookings))

	fmt.Printf("Thanks, %v %v for ordering %v tickets!\n", firstName, lastName, userTickets)
	fmt.Printf("You'll receive ticket at your email: %v\n\n", userEmail)
	fmt.Printf("Remaining tickets: %v\n", remTickets)

	firstNames := getFirstnames()
	fmt.Printf("All bookings: %v\n", firstNames)
}