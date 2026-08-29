package main

import "fmt"

func linearSearch(arr []int, v int) int {
	for i, item := range arr {
		if item == v {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{3, 7, 2, 9, 5, 1, 8, 4, 6}

	if idx := linearSearch(arr, 5); idx != -1 {
		fmt.Println(idx)
	} else {
		fmt.Println("Value not found.")
	}
}
