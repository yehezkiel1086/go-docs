package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Println("Can't read file")
	}
	defer file.Close() // this will execute last
}
