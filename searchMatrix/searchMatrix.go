package main

import "fmt"

func main() {

	fmt.Println(searchMatrix([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 5))
}

func searchMatrix(matrix [][]int, target int) bool {

	nums := []int{}
	for _, mt := range matrix {
		if mt[0] <= target && mt[len(mt)-1] >= target {
			nums = mt
			break
		}
	}
	for _, num := range nums {
		if target == num {
			return true
		}
	}

	return false
}
