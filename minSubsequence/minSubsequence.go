package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(minSubsequence([]int{4, 3, 10, 9, 8}))

}

func minSubsequence(nums []int) []int {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	afterSum, beforSum := 0, 0
	for _, num := range nums {
		afterSum += num
	}
	for i, num := range nums {
		beforSum += num
		afterSum -= num
		if beforSum > afterSum {
			return nums[:i+1]
		}
	}

	return nums
}
