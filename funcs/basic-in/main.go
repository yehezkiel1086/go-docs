package main

import "fmt"

func main() {
	var name, city, job, hobby string
	fmt.Print("Enter your name and city: ")
	fmt.Scan(&name, &city)

	fmt.Print("Enter your job and hobby: ")
	fmt.Scanf("%s %s", &job, &hobby)
	
	fmt.Printf("Welcome, %s! You are living in %s.\n", name, city)
	fmt.Printf("Your job is %s. You love %s.\n", job, hobby)
}