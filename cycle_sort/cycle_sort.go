package main

import (
	"fmt"
)

func cycleSort(arr []int) {
	n := len(arr)
	for cycleStart, item := range arr {
		pos := cycleStart
		for i := cycleStart + 1; i < n; i++ {
			if arr[i] < item {
				pos++
			}
		}
		if pos == cycleStart {
			continue
		}
		for item == arr[pos] {
			pos++
		}
		if pos != cycleStart {
			arr[pos], item = item, arr[pos]
		}
		for pos != cycleStart {
			pos = cycleStart
			for i := cycleStart + 1; i < n; i++ {
				if arr[i] < item {
					pos++
				}
			}
			for item == arr[pos] {
				pos++
			}
			if item != arr[pos] {
				arr[pos], item = item, arr[pos]
			}
		}
	}
}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	cycleSort(arr)
	fmt.Println(arr)
}
