package main

import (
	"fmt"
	"time"
)

var start time.Time

func concurrent() {
	apiCall := func(name string) {
		time.Sleep(time.Duration(100) * time.Millisecond)
		fmt.Println("Api Call", name, "running at", time.Since(start))
	}

	go apiCall("A")
	go apiCall("B")

	time.Sleep(time.Duration(200) * time.Millisecond)
	fmt.Println("Function ends at", time.Since(start))
}

func sequential() {
	apiCall := func(name string) {
		time.Sleep(time.Duration(100) * time.Millisecond)
		fmt.Println("Api Call", name, "running at", time.Since(start))
	}

	apiCall("A")
	apiCall("B")

	time.Sleep(time.Duration(200) * time.Millisecond)
	fmt.Println("Function ends at", time.Since(start))
}

func indecisive() {
	sayHello := func(name string) {
		fmt.Println("Hello", name)
	}

	go sayHello("Benjamin")
	go sayHello("Elisabeth")

	time.Sleep(time.Duration(10) * time.Millisecond)
}

func firstGoroutine() {
	go func() {
		fmt.Println("Goroutine thread")
	}()

	time.Sleep(time.Duration(10) * time.Millisecond)
}

func main() {
	fmt.Println("Main thread start")

	// firstGoroutine()
	// indecisive()

	// sequential vs concurrent
	start = time.Now()
	// concurrent()
	sequential()
	
	fmt.Println("Main thread end")
}
