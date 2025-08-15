package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Username string `json:"username"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Online    bool   `json:"online"`
}

func main() {
	user := User{
		Firstname: "Benjamin",
		Lastname:  "Franklin",
		Online:    false,
	}

	// convert object to json format
	jsonData, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))

	// read json file
	jsonContent, err := os.ReadFile("users.json")
	if err != nil {
		panic(err)
	}

	var users []User

	// convert json format to object
	if err := json.Unmarshal(jsonContent, &users); err != nil {
		panic(err)
	}

	fmt.Println(users)
}