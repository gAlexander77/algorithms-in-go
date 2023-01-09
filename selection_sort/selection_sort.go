package main

import (
	"fmt"
)

func selectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		// Find the minimum element in the unsorted part of the array.
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		// Swap the minimum element with the first element in the unsorted part.
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	selectionSort(arr)
	fmt.Println(arr) // [1, 2, 3, 4, 5]
}
