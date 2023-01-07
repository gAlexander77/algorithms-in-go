package main

import (
	"fmt"
)

func heapSort(arr []int) {
	// Build the heap.
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Extract elements from the heap one by one.
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func heapify(arr []int, n, i int) {
	root := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[root] {
		root = left
	}

	if right < n && arr[right] > arr[root] {
		root = right
	}

	if root != i {
		arr[i], arr[root] = arr[root], arr[i]
		heapify(arr, n, root)
	}
}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	heapSort(arr)
	fmt.Println(arr) // [1, 2, 3, 4, 5]
}
