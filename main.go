package main

import "fmt"

func main() {
	confName := "Go Conference";
	const confTickets uint = 50;
	var remTickets uint = 50;

	fmt.Printf("confTickets is %T, remTickets is %T", confTickets, remTickets);
	
	fmt.Printf("Welcome to %v booking app\n", confName);
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confTickets, remTickets);
	fmt.Println("Get your tickets here to attend");

	var firstName string;
	var lastName string;
	var email string;
	var userTickets uint;

	// ask user for their name
	fmt.Print("Enter your first name: ");
	fmt.Scan(&firstName);

	fmt.Print("Enter your last name: ");
	fmt.Scan(&lastName);

	fmt.Print("Enter your email address: ");
	fmt.Scan(&email);

	fmt.Print("Enter number of tickets: ");
	fmt.Scan(&userTickets);

	remTickets = remTickets - userTickets;

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email);
	fmt.Printf("%v tickets remaining for %v\n", remTickets, confName);
}