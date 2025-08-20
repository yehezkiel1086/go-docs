package main

import (
	"fmt"
	"time"
)

func printSomething(something string) {
	fmt.Println(something)
}

func simpleGoRoutine() {
	fmt.Println("Simple Go Routine:")
	fmt.Println("Function started")

	go printSomething("Routine called")

	fmt.Println("Before sleep")
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Function stopped")
}

func indecisiveRoutines() {
	fmt.Println("Indecisive Routines:")
	fmt.Println("Function started")

	go printSomething("1st routine called")
	go printSomething("2nd routine called")

	time.Sleep(10 * time.Millisecond)
	fmt.Println("Function stopped")
}

var start time.Time

func init() {
	start = time.Now()
}

func apiCallA() {
	time.Sleep(time.Millisecond * 100)
	fmt.Println("API Call A runs at:", time.Since(start))
}

func apiCallB() {
	time.Sleep(time.Millisecond * 100)
	fmt.Println("API Call B runs at:", time.Since(start))
}

func sequential() {
	fmt.Println("Sequential:")
	fmt.Println("Start time:", time.Since(start))
	apiCallA()
	apiCallB()

	time.Sleep(time.Millisecond * 200)
	fmt.Println("End time:", time.Since(start))
}

func concurrent() {
	fmt.Println("Concurrent:")
	fmt.Println("Start time:", time.Since(start))
	go apiCallA()
	go apiCallB()

	time.Sleep(time.Millisecond * 200)
	fmt.Println("End time:", time.Since(start))
}

func sequentialVsConcurrent() {
	sequential()
	fmt.Println()
	concurrent()
}

func main() {
	simpleGoRoutine()
	fmt.Println()
	indecisiveRoutines()
	fmt.Println()
	sequentialVsConcurrent()
}
