package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(countByOne(2))
	
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Can't read file")
	}

	defer file.Close()
}

func countByOne(x int) int {
    defer fmt.Println("Counting done")

    fmt.Println("Counting start")

    return x + 1
}
