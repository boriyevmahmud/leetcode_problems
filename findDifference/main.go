package main

import "fmt"

func main() {

	fmt.Println(findDifference([]int{1, 2, 3, 3}, []int{1, 1, 2, 2}))
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	resp := make([][]int, 2)
	numb0Map := make(map[int]bool)
	numb1Map := make(map[int]bool)
	for _, v1 := range nums1 {
		tmpBool := false
		for _, v2 := range nums2 {
			if v1 == v2 {
				tmpBool = true
			}
		}
		if !tmpBool && !numb0Map[v1] {
			numb0Map[v1] = true
			resp[0] = append(resp[0], v1)
		}
	}

	for _, v2 := range nums2 {
		tmpBool := false
		for _, v1 := range nums1 {
			if v2 == v1 {
				tmpBool = true
			}
		}
		if !tmpBool && !numb1Map[v2] {
			numb1Map[v2] = true

			resp[1] = append(resp[1], v2)
		}
	}

	return resp
}
