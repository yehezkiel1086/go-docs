package main

import (
	"fmt"
	"runtime"
)

func print(n int, message string) {
	for i := 0; i < n; i++ {
		fmt.Println((i + 1), message)
	}
}
func main() {
	runtime.GOMAXPROCS(2)

	go print(5, "Hello")
	print(5, "How are you?")

	var input string
	fmt.Scanln(&input)
}
