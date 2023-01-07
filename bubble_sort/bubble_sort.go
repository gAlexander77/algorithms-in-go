package main

import (
	"fmt"
)

func bubbleSort(arr []int) {
	// Set the flag to true to start the loop.
	flag := true

	// Loop until the flag is false.
	for flag {
		// Set the flag to false.
		flag = false

		// Loop through the array.
		for i := 0; i < len(arr)-1; i++ {
			// If the current element is greater than the next element, swap them and set the flag to true.
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				flag = true
			}
		}
	}
}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	bubbleSort(arr)
	fmt.Println(arr) // [1, 2, 3, 4, 5]

}
