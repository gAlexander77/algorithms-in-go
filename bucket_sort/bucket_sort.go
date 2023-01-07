package main

import (
	"fmt"
)

func bucketSort(arr []int, bucketSize int) []int {
	// Get the minimum and maximum values in the array.
	minValue := arr[0]
	maxValue := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < minValue {
			minValue = arr[i]
		} else if arr[i] > maxValue {
			maxValue = arr[i]
		}
	}

	// Calculate the number of buckets.
	bucketCount := (maxValue-minValue)/bucketSize + 1

	// Create a slice of slices to hold the buckets.
	buckets := make([][]int, bucketCount)
	for i := 0; i < bucketCount; i++ {
		buckets[i] = make([]int, 0)
	}

	// Put the elements into the buckets.
	for i := 0; i < len(arr); i++ {
		bucketIndex := (arr[i] - minValue) / bucketSize
		buckets[bucketIndex] = append(buckets[bucketIndex], arr[i])
	}

	// Sort the elements in each bucket.
	for i := 0; i < bucketCount; i++ {
		insertionSort(buckets[i])
	}

	// Concatenate the buckets into a single slice.
	result := make([]int, 0)
	for i := 0; i < bucketCount; i++ {
		result = append(result, buckets[i]...)
	}
	return result
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}
}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	sortedArr := bucketSort(arr, 1)
	fmt.Println(sortedArr) // [1, 2, 3, 4, 5]

}
