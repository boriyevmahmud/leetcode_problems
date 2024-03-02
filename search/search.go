package main

import "fmt"

func main() {

	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}

func search(nums []int, target int) int {

	for i, num := range nums {
		if num == target {
			return i
		}
	}

	return -1
}
