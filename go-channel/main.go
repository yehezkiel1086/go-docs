package main

import (
	"fmt"
	"runtime"
	"time"
)

// simple channel
func simpleChan() {
	msg := make(chan string)

	sayHelloTo := func(name string) {
		data := fmt.Sprintf("Hello, %s", name)
		msg <- data
	}

	names := []string{"Yehezkiel", "Ben", "Yehuda"}

	for i := range len(names) {
		go sayHelloTo(names[i])
		fmt.Println(<-msg)
	}
}

// buffered channel
func bufferedChan() {
	msg := make(chan string, 3)

	// receive
	go func() {
		for {
			data := <- msg
			fmt.Println(data)
		}
	}()

	// send
	for _, name := range []string{"Ben", "Yehuda", "Yehezkiel", "Tzivah", "Yehochanan"} {
		data := fmt.Sprintf("Hello, %s", name)
		msg <- data
	}

	time.Sleep(200 * time.Millisecond)
}

func main() {
	runtime.GOMAXPROCS(2)

	// simpleChan()
	bufferedChan()
}
