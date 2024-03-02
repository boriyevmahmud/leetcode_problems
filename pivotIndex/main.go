package main

import "fmt"

func main() {

	fmt.Println(pivotIndex([]int{1, 2, 3}))

}

func pivotIndex(nums []int) int {
	totalSum := 0

	for _, num := range nums {
		totalSum += num
	}

	leftSum := 0

	for i := 0; i < len(nums); i++ {
		if leftSum == totalSum-leftSum-nums[i] {
			return i
		}

		leftSum += nums[i]
	}

	return -1
}
