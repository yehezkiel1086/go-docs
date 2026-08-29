package main

import (
	"fmt"
	"runtime"
	"sync"
)

type Counter struct {
	mtx sync.Mutex
	val int
}

func (c *Counter) Add() {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.val++
}

func (c *Counter) Value() int {
	return c.val
}

func main() {
	runtime.GOMAXPROCS(2)

	var c Counter
	var wg sync.WaitGroup

	for range 1000 {
		wg.Add(1)

		go func() {
			for range 1000 {
				c.Add()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(c.Value())
}
