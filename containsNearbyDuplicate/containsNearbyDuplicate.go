package main

import (
	"fmt"
)

func main() {

	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	tmpMap := make(map[int]int, 0)
	for i, num := range nums {
		if v, ok := tmpMap[num]; ok && i-v <= k {
			return true
		}
		tmpMap[num] = i
	}
	return false
}
