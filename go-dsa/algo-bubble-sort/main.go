package main

import "fmt"

/**
1. Go through the array, one value at a time.
2. For each value, compare the value with the next value.
3. If the value is higher than the next one, swap the values so that the highest value comes last.
4. Go through the array as many times as there are values in the array.
*/

func bubbleSort(arr []int) []int {
	for i := range arr {
		for j := 0; j < len(arr) - i - 1; j += 1 {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90, 5}
	sorted := bubbleSort(arr)
	fmt.Println(sorted)
}
