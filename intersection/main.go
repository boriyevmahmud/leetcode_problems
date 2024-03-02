package main

import "fmt"

func main() {

	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}

func intersection(nums1 []int, nums2 []int) []int {
	numsMap := make(map[int]bool)
	result := make([]int,0, len(nums1)+len(nums2))

	for _, n := range nums1 {
		numsMap[n] = true
	}

	for _, n := range nums2 {
		if numsMap[n] {
			result = append(result, n)
			numsMap[n] = false
		}

	}

	return result
}
