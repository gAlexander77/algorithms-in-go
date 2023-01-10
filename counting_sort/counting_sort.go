package main

import (
	"fmt"
)

func countingSort(arr []int) []int {
	maxValue := arr[0]
	for _, v := range arr {
		if v > maxValue {
			maxValue = v
		}
	}

	counts := make([]int, maxValue+1)
	for _, v := range arr {
		counts[v]++
	}

	for i := 1; i <= maxValue; i++ {
		counts[i] += counts[i-1]
	}

	sortedArr := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		sortedArr[counts[arr[i]]-1] = arr[i]
		counts[arr[i]]--
	}
	return sortedArr
}

func main() {
	arr := []int{4, 2, 9, 6, 23, 12, 34, 0, 1}
	fmt.Println(countingSort(arr))
}
