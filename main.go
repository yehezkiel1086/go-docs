package main

import "fmt"

func main() {
	var confName = "Go Conference";
	const confTickets = 50;
	var remTickets = 50;
	
	fmt.Printf("Welcome to %v booking app\n", confName);
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confTickets, remTickets);
	fmt.Println("Get your tickets here to attend");
}