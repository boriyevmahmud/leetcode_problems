package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(increasingTriplet([]int{20, 100, 10, 12, 5, 13}))
}

// func increasingTriplet(nums []int) bool {
// 	if len(nums) <= 2 {
// 		return false
// 	}

// 	for i := 0; i < len(nums)-2; i++ {
// 		for j := i; j < len(nums)-1; j++ {
// 			for k := j; k < len(nums); k++ {
// 				if nums[i] < nums[j] && nums[j] < nums[k] {
// 					return true
// 				}
// 			}
// 		}
// 	}

// 	return false
// }

func increasingTriplet(nums []int) bool {
	if len(nums) <= 2 {
		return false
	}

	firstMin, secondMin := math.MaxInt32, math.MaxInt32

	for _, num := range nums {
		if num <= firstMin {
			firstMin = num
		} else if num <= secondMin {
			secondMin = num
		} else {
			// We found a number greater than firstMin and secondMin
			return true
		}
	}

	return false
}
