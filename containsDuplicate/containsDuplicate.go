package main

import "fmt"

func main() {

	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
}

func containsDuplicate(nums []int) bool {
	tmpMap := make(map[int]bool)
	for _, num := range nums {
		if ok := tmpMap[num]; ok {
			return true
		}
		tmpMap[num] = true
	}
	return false
}
