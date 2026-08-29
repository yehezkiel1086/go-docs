package main

import "fmt"

/**
1. Check the value in the center of the array.
2. If the target value is lower, search the left half of the array. If the target value is higher, search the right half.
3. Continue step 1 and 2 for the new reduced part of the array until the target value is found or until the search area is empty.
4. If the value is found, return the target value index. If the target value is not found, return -1.
*/

func binarySearch(arr []int, v int) int {
	l := 0
	r := len(arr) - 1

	for l <= r {
		m := (l + r) / 2

		if v == arr[m] {
			return m
		} else if arr[m] < v {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

	if idx := binarySearch(arr, 11); idx != -1 {
		fmt.Println(idx)
	} else {
		fmt.Println("Value not found.")
	}
}
