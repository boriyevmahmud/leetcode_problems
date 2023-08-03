package main

import "fmt"

func linearSearch(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1 // target not found
}

func main() {
	arr := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	target := 23

	result := linearSearch(arr, target)

	if result == -1 {
		fmt.Printf("%d not found in the array\n", target)
	} else {
		fmt.Printf("%d found at index %d\n", target, result)
	}
}
