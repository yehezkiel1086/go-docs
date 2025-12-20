package main

import (
	"fmt"
	"time"
)

func send(c chan int) {
	for i := range(3) {
		fmt.Printf("send: %d\n", i)
		c <- i
	}
}

func capacityBlocking() {
	c := make(chan int, 1)

	go send(c)

	for range(3) {
		fmt.Printf("Receive: %d\n", <-c)
	}
}

func infLoop() {
	c := make(chan string, 1)

	go func() {
		for range(10) {
			c <- "every 200 ms"
			time.Sleep(200 * time.Millisecond)
		}
		close(c) // prevents deadlock
	}()

	for msg := range c {
		fmt.Println(msg)
	}
}

func infLoopClose() {
	c := make(chan string, 5)

	for i := range(5) {
		c <- fmt.Sprintf("data %d", i)
	}

	close(c) // prevents deadlock

	for msg := range c {
		fmt.Println(msg)
	}
}

var start time.Time

func init() {
	start = time.Now()
}

func multiChannels() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			c1 <- "every 200 milliseconds"
		}
	}()

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			c2 <- "every 500 milliseconds"
		}
	}()

	for {
		select {
		case msg := <-c1:
			fmt.Printf("%s channel 1. At time %s\n", msg, time.Since(start))
		case msg := <-c2:
			fmt.Printf("%s channel 2. At time %s\n", msg, time.Since(start))
		}
	}
}

func main() {
	// capacityBlocking()
	// infLoop()
	// infLoopClose()
	multiChannels()
}
