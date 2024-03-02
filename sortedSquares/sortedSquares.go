package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}

func sortedSquares(nums []int) []int {

	squaredArr := make([]int, 0, len(nums))
	for _, num := range nums {
		squaredArr = append(squaredArr, num*num)
	}
	sort.Ints(squaredArr)
	return squaredArr
}
