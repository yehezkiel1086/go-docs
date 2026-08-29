package util

import (
	"fmt"
	"time"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

func APICall() {
	fmt.Println("main thread started")

	// sequential
	// APICallA()
	// APICallB()

	// concurrent
	go APICallA()
	go APICallB()

	time.Sleep(400 * time.Millisecond)
	fmt.Println("main thread ends at", time.Since(startTime))
}

func APICallA() {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("API Call A ends at", time.Since(startTime))
}

func APICallB() {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("API Call B ends at", time.Since(startTime))
}
