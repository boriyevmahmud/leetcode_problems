package main

import (
	"fmt"
	"sort"
)

func main() {

	// fmt.Println(targetIndices([]int{75, 99, 19, 93, 87, 68, 12, 18, 48, 83, 24, 50, 16, 53, 36, 16, 80, 68, 46, 13, 53, 100, 50, 49, 77, 52, 34, 42, 38, 98, 73, 11, 13, 61, 72, 8, 11, 67, 98, 24, 23, 71, 47, 6, 5, 7, 97, 86, 25, 82, 11, 15, 26, 97, 69, 6, 30, 77, 98, 44, 32, 39, 71, 47, 64, 78, 6, 61, 72, 75}, 98))
	fmt.Println(targetIndices([]int{1, 2, 5, 2, 3}, 5))
}

func targetIndices(nums []int, target int) []int {

	sort.Ints(nums)

	return searchRange(nums, target)
}

// func searchRange(nums []int, target int) []int {

// 	resp := make([]int, 0)

// 	for i, num := range nums {
// 		if i == len(nums)-1 {
// 			break
// 		}
// 		if num == target && len(resp) == 0 {
// 			resp = append(resp, i)
// 		} else if num == target && nums[i+1] > target && len(resp) == 1 {
// 			resp = append(resp, i)
// 		}
// 	}

// 	if nums[len(nums)-1] == target {
// 		resp = append(resp, len(nums)-1)
// 	}

// 	return resp
// }
func searchRange(nums []int, target int) []int {
	resp := make([]int, 0)

	fmt.Println("Searching", nums)
	for i, num := range nums {
		if num == target {
			resp = append(resp, i)
		}
	}

	return resp
}
