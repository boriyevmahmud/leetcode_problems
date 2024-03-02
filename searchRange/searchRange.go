package main

import "fmt"

func main() {

	fmt.Println(searchRange([]int{2, 2}, 2))
}
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	} else if len(nums) == 1 && nums[0] == target {
		return []int{0, 0}
	}

	resp := make([]int, 0)

	for i, num := range nums {
		if i == len(nums)-1 {
			break
		}
		if num == target && len(resp) == 0 {
			resp = append(resp, i)
		} else if num == target && nums[i+1] > target && len(resp) == 1 {
			resp = append(resp, i)
		}
	}

	if nums[len(nums)-1] == target {
		resp = append(resp, len(nums)-1)
	}

	if len(resp) == 0 {
		return []int{-1, -1}
	} else if len(resp) == 1 {
		resp = append(resp, resp[0])
		return resp
	}

	return resp
}
