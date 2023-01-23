package main

import (
	"fmt"
)

func combSort(arr []int) {
	n := len(arr)
	gap := n
	shrink := 1.3
	sorted := false
	for !sorted {
		gap = int(float64(gap) / shrink)
		if gap > 1 {
			sorted = false
		} else {
			gap = 1
			sorted = true
		}
		i := 0
		for i+gap < n {
			if arr[i] > arr[i+gap] {
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
				sorted = false
			}
			i++
		}
	}
}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	combSort(arr)
	fmt.Println(arr)
}
