package main

import "fmt"

func main() {
	fmt.Println(countPairs([]int{-1, 1, 2, 3, 1}, 2))
}

func countPairs(nums []int, target int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i]+nums[j] < target && (0 <= i && i < j && j < len(nums)) {
				cnt += 1
			}
		}
	}
	return cnt
}
