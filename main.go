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

	var username string;
	var userTickets uint;

	// ask user for their name

	username = "Ben";
	userTickets = 2;
	fmt.Printf("User %v booked %v tickets.\n", username, userTickets);
}