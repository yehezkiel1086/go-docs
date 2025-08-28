package main

import "fmt"

func sender(c chan int) {
	for i := 0; i < 4; i++ {
		fmt.Println("Goroutine send:", i)
		c <- i // sends 4 times
	}
}

func receiver() {
	c := make(chan int, 1) // size of buffer channel just 1
	go sender(c)
	for i := 0; i < 3; i++ {
		fmt.Println("receiver receive:", <- c) // receives three times
	}
}

func main() {
		receiver()
}
