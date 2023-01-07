package main

import (
	"fmt"
)

func linearSearch(arr []int, target int) int {
	// Loop through the array and return the index if the target is found.
	for i, value := range arr {
		if value == target {
			return i
		}
	}

	// Return -1 if the target was not found.
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	target := 3
	index := linearSearch(arr, target)
	fmt.Println(index) // 2

}
