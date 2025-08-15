package main

import (
	"fmt"
	"strings"
)

func main() {
	nums := []float64{1, 5, 2, 4, 3, 8, 6, 9, 7, 10}
	res := calcAvg(nums...)
	fmt.Printf("Average: %v\n", res)

	printHobbies("Yehezkiel", "CP", "CTF", "OffSec", "ML")
}

// variadic funcs
func calcAvg(nums ...float64) float64 {
	tt := 0.0

	for _, x := range nums {
		tt += x
	}

	return tt / float64(len(nums))
}

func printHobbies(name string, hobbies ...string) {
	fmt.Printf("%v hobbies: %v\n", name, strings.Join(hobbies, ", "))
}
