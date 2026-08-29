package main

import (
	"fmt"
	"runtime"
	"time"
)

// channel as parameter
func sayHello(msg chan string) {
	fmt.Println(<-msg)
}

func printMessages() {
	msg := make(chan string)

	setMsg := func(name string) {
		data := fmt.Sprintf("Hello, %v", name)
		msg <- data
	}

	go setMsg("Benjamin")
	go setMsg("Elizabeth")
	go setMsg("Maria")

	sayHello(msg)
	sayHello(msg)
	sayHello(msg)
}

func chanSync() {
	runtime.GOMAXPROCS(1)
	start := time.Now()

	fmt.Println("Main started at time", time.Since(start))
	c := make(chan string)
	go func() {
		time.Sleep(time.Duration(10) * time.Millisecond)
		fmt.Println("Hello from goroutine at time", time.Since(start))
		c <- "Goroutine say hi"
	}()

	fmt.Printf("Goroutine sent this: %v, at time %v\n", <-c, time.Since(start))
	fmt.Println("Main ended at time", time.Since(start))
}

func main() {
	// printMessages()
	chanSync()
}
