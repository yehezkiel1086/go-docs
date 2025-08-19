package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func main() {
	fmt.Println(math.Abs(3.4))
	fmt.Println(math.Floor(3.4))
	fmt.Println(math.Ceil(3.4))

	fmt.Println(math.Pow(2, 3))

	// random numbers
	/*
		if you don't set the seed value, if you print it again and again, the value will stay the same (it's required in newer versions of Go)
	*/
	// rand.Seed(10)
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println(rand.Int())
	fmt.Println(rand.Int())
	fmt.Println(rand.Int())

	// random with custom range
	fmt.Println("0 - 20:", rand.Intn(20))
	fmt.Println("0 - 100:", rand.Intn(100))

	fmt.Println(randStr(10))
}
