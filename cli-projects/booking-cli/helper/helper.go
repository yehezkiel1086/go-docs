package helper

import (
	"fmt"
	"strings"
)

func ValidateInput(firstName string, lastName string, userEmail string, userTickets uint8, remTickets uint8) bool {
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