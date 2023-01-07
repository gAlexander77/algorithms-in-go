package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) int {
	// Set the start and end indices.
	start := 0
	end := len(arr) - 1

	for start <= end {
		// Calculate the midpoint.
		mid := (start + end) / 2

		// Check if the target is less than, greater than, or equal to the midpoint.
		if target < arr[mid] {
			end = mid - 1
		} else if target > arr[mid] {
			start = mid + 1
		} else {
			return mid
		}
	}

	// If the target was not found, return -1.
	return -1
}
func main() {
	arr := []int{1, 2, 3, 4, 5}
	target := 3
	index := binarySearch(arr, target)
	fmt.Println(index) // 2
}
