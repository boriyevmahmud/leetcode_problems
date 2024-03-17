package main

import "fmt"

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))

}

func findDisappearedNumbers(nums []int) []int {

	numMap := make(map[int]bool)
	maxNum := nums[0]
	resp := []int{}

	for _, num := range nums {
		numMap[num] = true
		if num > maxNum {
			maxNum = num
		}
	}

	for i := 1; i <= len(nums); i++ {
		if !numMap[i] {
			resp = append(resp, i)
		}
	}

	return resp
}
