package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
	fmt.Println(productExceptSelf([]int{1, 0}))
	fmt.Println(productExceptSelf([]int{0, 0}))
}

func productExceptSelf(nums []int) []int {
	res := make([]int, 0, len(nums))
	tmpNumb := 1
	zeroUsed := false
	for _, n := range nums {
		if n == 0 {
			zeroUsed = true
			continue
		}
		tmpNumb *= n
	}
	for _, n := range nums {
		if n == 0 && tmpNumb != 1 {
			res = append(res, tmpNumb)
		} else if !zeroUsed {
			res = append(res, tmpNumb/n)
		} else {
			res = append(res, 0)

		}
	}

	return res
}
