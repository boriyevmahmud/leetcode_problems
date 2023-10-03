package main

import "fmt"

func main() {

	fmt.Println(removeDuplicates([]int{1, 1, 2, 2, 2, 3, 4}))
}

func removeDuplicates(nums []int) int {
	prev := nums[0]
	l := 1
	for i := 1; i < len(nums); i++ {
		if prev != nums[i] {
			nums[l] = nums[i]
			l++
		}
		prev = nums[i]
	}
	fmt.Println(nums)
	return l
}
