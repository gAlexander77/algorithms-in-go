package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]

	less := make([]int, 0)
	greater := make([]int, 0)

	for i, value := range arr {
		if i == pivotIndex {
			continue
		}
		if value < pivot {
			less = append(less, value)
		} else {
			greater = append(greater, value)
		}
	}

	return append(append(quickSort(less), pivot), quickSort(greater)...)
}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	sortedArr := quickSort(arr)
	fmt.Println(sortedArr)
}
