package main

import (
	"fmt"
)

func mergeSort(arr []int) []int {
	// If the array has less than 2 elements, return it.
	if len(arr) < 2 {
		return arr
	}

	// Split the array in half.
	middle := len(arr) / 2
	left := arr[:middle]
	right := arr[middle:]

	// Recursively sort the left and right halves.
	left = mergeSort(left)
	right = mergeSort(right)

	// Merge the sorted halves.
	return merge(left, right)
}

func merge(left, right []int) []int {
	// Create a new slice to hold the merged elements.
	result := make([]int, 0)

	// Set the indices for the left and right slices.
	leftIndex := 0
	rightIndex := 0

	// Loop until one of the slices is exhausted.
	for leftIndex < len(left) && rightIndex < len(right) {
		// If the left element is smaller, append it to the result and increment the left index.
		if left[leftIndex] < right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			// Otherwise, append the right element and increment the right index.
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}

	// Append the remaining elements from the left slice, if any.
	result = append(result, left[leftIndex:]...)

	// Append the remaining elements from the right slice, if any.
	result = append(result, right[rightIndex:]...)

	return result
}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	sortedArr := mergeSort(arr)
	fmt.Println(sortedArr) // [1, 2, 3, 4, 5]

}
