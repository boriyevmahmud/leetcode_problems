package main

import "fmt"

func main() {

	moveZeroes([]int{1, 0, 0, 2, 3})
}

// func moveZeroes(nums []int) {
// 	nonZeroIndex := 0
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] == 0 {
// 			if nums[i] == 0 {
// 				nums[nonZeroIndex], nums[i] = nums[nonZeroIndex], nums[i]
// 				nonZeroIndex++
// 			}
// 		}
// 	}
// 	fmt.Println("nums: ", nums)

// }

func moveZeroes(nums []int) {
	nonZeroIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[nonZeroIndex], nums[i] = nums[i], nums[nonZeroIndex]
			nonZeroIndex++
		}
	}
}
