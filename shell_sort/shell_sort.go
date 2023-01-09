package main

import (
	"fmt"
)

func shellSort(arr []int) {
	// Set the gap between elements.
	gap := len(arr) / 2
	for gap > 0 {
		// Insertion sort on the gap-separated elements.
		for i := gap; i < len(arr); i++ {
			j := i
			for j >= gap && arr[j-gap] > arr[j] {
				arr[j], arr[j-gap] = arr[j-gap], arr[j]
				j -= gap
			}
		}

		// Decrease the gap.
		gap /= 2
	}
}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	shellSort(arr)
	fmt.Println(arr) // [1, 2, 3, 4, 5]

}
