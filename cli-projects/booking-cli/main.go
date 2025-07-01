package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var confName string
const confTickets uint8 = 50
var remTickets uint8 = 50
var bookings = make([]UserData, 0)
// slice
// var bookings []string

type UserData struct {
	firstName string
	lastName string
	email string
	tickets uint8
}

var wg = sync.WaitGroup{}

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

	firstName, lastName, userEmail, userTickets := getUserInput()

	isValidInput := helper.ValidateInput(firstName, lastName, userEmail, userTickets, remTickets)

	if isValidInput {
		bookTicket(userTickets, firstName, lastName, userEmail)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, userEmail)

		if remTickets == 0 {
			fmt.Println("Ticket is sold out. Come back next year!")
			// break
		}
	}

	wg.Wait()
}

func greetUsers() {
	fmt.Println("Welcome to", confName, "🥳")
	fmt.Println("Note: Get your tickets here to attend.")
	fmt.Printf("There are %v left out of %v\n", remTickets, confTickets)
}

func getFirstnames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
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
	// create a map for user
	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: userEmail,
		tickets: userTickets,
	}

	bookings = append(bookings, userData) // add element to slice
	remTickets -= userTickets
	fmt.Printf("All bookings: %v\n", bookings)
	// bookings[0] = userName // add element to array

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

func sendTicket(tickets uint8, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", tickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n%v to email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}
