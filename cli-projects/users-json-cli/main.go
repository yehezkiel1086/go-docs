package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Online    bool   `json:"online"`
}

func main() {
	users := getUsersData()
	user := createUser("Itamar", "Ben Gvir", true)

	users = append(users, user)

	jsonData, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("users.json", jsonData, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))
}

func createUser(firstName string, lastName string, online bool) User {
	return User{
		FirstName: firstName,
		LastName: lastName,
		Online: online,
	}
}

func getUsersData() []User {
	var users []User
	content, err := os.ReadFile("users.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(content, &users)
	if err != nil {
		panic(err)
	}

	return users
}
