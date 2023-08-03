package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]

	left := make([]int, 0)
	right := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		if i == pivotIndex {
			continue
		}
		if arr[i] <= pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, pivot), right...)
}

func main() {
	arr := []int{64, 25, 12, 22, 11}
	sortedArr := quickSort(arr)
	fmt.Println(sortedArr)
}
