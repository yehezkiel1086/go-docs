package main

import (
	"fmt"
	"sync"
)

func doPrint(wg *sync.WaitGroup, msg string) {
	defer wg.Done()
	fmt.Println(msg)
}

func main() {
	var wg sync.WaitGroup

	for i := range 5 {
		msg := fmt.Sprintf("Current: %d", i)

		wg.Add(1)
		go doPrint(&wg, msg)
	}

	wg.Wait()
}
