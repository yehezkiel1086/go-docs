package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var jsonData = `{ "first_name" : "Sammy", "last_name": "Shark",  "online" : true }`

type User struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Online bool `json:"online"`
}

func main() {
	var user1, user2 User
	err := json.Unmarshal([]byte(jsonData), &user1)
	if err != nil {
		panic(err)
	}

	user2 = User{
		FirstName: "Benjamin",
		LastName: "Meir",
		Online: false,
	}
	user2Json, err := json.Marshal(user2)

	if err != nil {
		panic(err)
	}

	fmt.Println(user1, string(user2Json))

	var user3 User
	jsonData, err := os.ReadFile("user.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(jsonData), &user3)

	if err != nil {
		panic(err)
	}

	fmt.Println(user3)

	jsonDatas := [3]User{user1, user2, user3}

	dataStr, err := json.Marshal(jsonDatas)

	if err != nil {
		panic(err)
	}	

	err = os.WriteFile("users.json", dataStr, 0644)

	if err != nil {
		panic(err)
	}
}
