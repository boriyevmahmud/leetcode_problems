package main

import "fmt"

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1 // target not found
}

func main() {
	arr := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	target := 23

	result := binarySearch(arr, target)

	if result == -1 {
		fmt.Printf("%d not found in the array\n", target)
	} else {
		fmt.Printf("%d found at index %d\n", target, result)
	}
}
