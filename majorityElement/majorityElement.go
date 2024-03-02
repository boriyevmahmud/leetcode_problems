package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{3, 2, 2, 2,  1, 1, 2, 3}))

}

//optimized version
func majorityElement(nums []int) int {
	majority_elem, majority_frq := nums[0], 1

	for i := 1; i < len(nums); i++ {
		if majority_frq == 0 {
			majority_elem, majority_frq = nums[i], 1
		} else {

			if nums[i] == majority_elem {
				majority_frq += 1
			} else {
				majority_frq -= 1
			}
		}
	}

	return majority_elem
}

// first solution
// func majorityElement(nums []int) int {

// 	numbs := make(map[int]int)

// 	for _, val := range nums {
// 		numbs[val] += 1
// 	}

// 	max := 0
// 	res := 0
// 	for key, v := range numbs {
// 		if v > max {
// 			max = v
// 			res = key
// 		}
// 	}

// 	return res
// }
