package main

import "fmt"

func main() {

	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 3}))
}

// func removeDuplicates(nums []int) int {
// 	prev := nums[0]
// 	l := 2
// 	for i := 2; i < len(nums); i+=2 {
// 		if prev != nums[i] {
// 			nums[l] = nums[i]
// 			l+=2
// 		}
// 		prev = nums[i]
// 	}
// 	fmt.Println(nums)
// 	return l
// }

func removeDuplicates(nums []int) int {

	if len(nums) < 2 {
		return len(nums)
	}

	idx := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[idx-2] {
			nums[idx] = nums[i]
			idx += 1
		}
	}

	return idx
}
