package util

import (
	"fmt"
	"time"
)

func Hello() {
	fmt.Println("main thread started")

	sayHello := func(name string) {
		fmt.Println("Hello,", name)
	}

	go sayHello("John")
	go sayHello("Jane")
	go sayHello("Yehezkiel")

	time.Sleep(200 * time.Millisecond)
	fmt.Println("main thread ends at", time.Since(startTime))
}
