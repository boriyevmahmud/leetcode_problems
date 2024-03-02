package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(maximumGap([]int{3, 6, 9, 15}))

}

func maximumGap(nums []int) int {

	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)

	max := 0
	for i := 0; i < len(nums)-1; i++ {
		if max < nums[i+1]-nums[i] {
			max = nums[i+1] - nums[i]
		}
	}

	return max
}
