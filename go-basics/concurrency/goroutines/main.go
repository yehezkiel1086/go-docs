package main

import (
	"fmt"
	"time"
)

func simpleRoutine() {
	fmt.Println("Simple Routine:")

	go func() {
		fmt.Println("Simple Routine")
	}()

	fmt.Println("Routine start")
	time.Sleep(time.Millisecond * 10)
	fmt.Println("Routine end")
}

func indecisiveRoutine() {
	fmt.Println("Indecisive Routine:")

	go func() {
		fmt.Println("Routine A")
	}()
	go func() {
		fmt.Println("Routine B")
	}()

	fmt.Println("Routine started")
	time.Sleep(time.Millisecond * 10)
	fmt.Println("Routine end")
}

func ApiCallA(start time.Time) {
	time.Sleep(time.Millisecond * 100)
	fmt.Println("API call A at:", time.Since(start))
}

func ApiCallB(start time.Time) {
	time.Sleep(time.Millisecond * 100)
	fmt.Println("API call B at:", time.Since(start))
}

func sequentialApiCall() {
	fmt.Println("Sequential API call Begins:")
	start := time.Now()
	ApiCallA(start)
	ApiCallB(start)

	time.Sleep(time.Millisecond * 200)
	fmt.Println("Sequential API call ends at:", time.Since(start))
}

func concurrentApiCall() {
	fmt.Println("Concurrent API call Begins:")
	start := time.Now()
	go ApiCallA(start)
	go ApiCallB(start)

	time.Sleep(time.Millisecond * 200)
	fmt.Println("Sequential API call ends at:", time.Since(start))
}

// from APICallA at time 111.1639ms
// from APICallB at time 218.4708ms
// from main function at time 420.4578ms

func main() {
	simpleRoutine()
	fmt.Println()
	indecisiveRoutine()
	fmt.Println()
	sequentialApiCall()
	fmt.Println()
	concurrentApiCall()
}
