package main

import "fmt"

func main() {
	var confName = "Go Conference";
	const confTickets = 50;
	var remTickets = 50;
	
	fmt.Println("Welcome to", confName, "booking app");
	fmt.Println("We have total of", confTickets, "ticketes and", remTickets, "are still available.");
	fmt.Println("Get your tickets here to attend");
}