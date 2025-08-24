package main

import "fmt"

func sayHelloTo(name string, msgs chan string) {
	msgs <- fmt.Sprintf("Hello, %v", name)
}

func printMessage(msgs chan string) {
	fmt.Println(<-msgs)
}

func sayHelloChannels() {
	names := []string{"Ben", "Andre", "Yehuda"}
	msgs := make(chan string)

	for _, item := range names {
		go sayHelloTo(item, msgs)
	}

	for i := 0; i < len(names); i++ {
		printMessage(msgs)
	}
}

func main() {
	sayHelloChannels()
}