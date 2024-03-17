package main

import "fmt"

func main() {
	fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

func findDuplicates(nums []int) []int {
	if len(nums) == 1 {
		return nil
	}

	resp := make([]int, 0, len(nums))
	numbersMap := make(map[int]int)

	for _, n := range nums {
		if v, ok := numbersMap[n]; ok && v == 1 {
			resp = append(resp, n)
		}
		numbersMap[n]++

	}

	return resp
}
