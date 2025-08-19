package main

import "fmt"

// function as parameter
func apply(f func(uint8) uint8, x uint8) uint8 {
	return f(x)
}

func main() {
	// function as variable
	inc := func(val uint8) uint8 {
		return val + 1
	}
	dec := func(val uint8) uint8 {
		return val - 1
	}

	fmt.Println(inc(2))

	// apply function as parameter
	fmt.Println(apply(dec, 2))

	// IIFE: Immediately Invoked Function
	arr := []uint8{1, 2, 3, 4, 5}

	avg := func(lst []uint8) float32 {
		tt := 0
		for i := range arr {
			tt += i
		}
		return float32(tt) / float32(len(arr))
	}(arr)

	fmt.Printf("%.2f\n", avg)
}
