package main

import "fmt"

/**
  Create an empty list (it can also be a dictionary or a set).
  Create a hash function.
  Inserting an element using a hash function.
  Looking up an element using a hash function.
  Handling collisions.
*/

// create empty list
var users [10][]string

// create hash function
func hashFunc(v string) int {
	charsSum := 0
	for _, i := range v {
		charsSum += int(i)
	}
	return charsSum % 10
}

// insert elements using hash function
func add(v string) {
	idx := hashFunc(v)
	users[idx] = append(users[idx], v)
}

// lookup element
func contains(v string) bool {
	idx := hashFunc(v)
	for _, i := range users[idx] {
		if i == v {
			return true
		}
	}
	return false
}

func main() {
	add("Bob")
	add("Pete")
	add("Jones")
	add("Lisa")
	add("Stuart")
	add("Siri")
	fmt.Println(users)
	fmt.Println(contains("Pete"))
}
