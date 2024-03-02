package main

import "fmt"

func main() {

	fmt.Println(findMin([]int{1, 2, 3, 4, 5}))
}

func findMin(nums []int) int {
	min := nums[0]
	for _, num := range nums {
		if min > num {
			min = num
		}
	}
	return min
}
