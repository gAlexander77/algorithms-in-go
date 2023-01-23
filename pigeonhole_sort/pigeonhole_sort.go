package main

import (
	"fmt"
)

func pigeonholeSort(arr []int) {
	n := len(arr)
	// find the minimum and maximum values in the list
	min := arr[0]
	max := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}
	// create the pigeonholes
	pigeonholes := make([]int, max-min+1)
	// put the elements into the pigeonholes
	for i := 0; i < n; i++ {
		pigeonholes[arr[i]-min]++
	}
	// take the elements out of the pigeonholes
	index := 0
	for i := 0; i < len(pigeonholes); i++ {
		for pigeonholes[i] > 0 {
			arr[index] = i + min
			index++
			pigeonholes[i]--
		}
	}
}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	pigeonholeSort(arr)
	fmt.Println(arr)
}
