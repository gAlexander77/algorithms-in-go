package main

import (
	"fmt"
)

func radixSort(arr []int) []int {
	// Get the maximum number of digits.
	maxDigits := getMaxDigits(arr)

	// Loop through the digits, starting with the least significant.
	for digit := 1; digit <= maxDigits; digit++ {
		// Create the buckets for each digit.
		buckets := make([][]int, 10)
		for i := range buckets {
			buckets[i] = make([]int, 0)
		}

		// Distribute the elements into the buckets based on the current digit.
		for _, value := range arr {
			digitValue := getDigit(value, digit)
			buckets[digitValue] = append(buckets[digitValue], value)
		}

		// Concatenate the buckets into a single slice.
		arr = make([]int, 0)
		for _, bucket := range buckets {
			arr = append(arr, bucket...)
		}
	}

	return arr
}

func getMaxDigits(arr []int) int {
	maxDigits := 0
	for _, value := range arr {
		digits := getNumDigits(value)
		if digits > maxDigits {
			maxDigits = digits
		}
	}
	return maxDigits
}

func getNumDigits(value int) int {
	count := 0
	for value > 0 {
		value /= 10
		count++
	}
	return count
}

func getDigit(value, digit int) int {
	for i := 1; i < digit; i++ {
		value /= 10
	}
	return value % 10
}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	sortedArr := radixSort(arr)
	fmt.Println(sortedArr) // [1, 2, 3, 4, 5]
}
