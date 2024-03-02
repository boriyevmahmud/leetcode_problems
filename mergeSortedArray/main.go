package main

import (
	"sort"
)

func main() {

	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	if len(nums1) > m {
		nums1 = nums1[:m]
	}
	if len(nums2) > n {
		nums2 = nums2[:n]
	}
	for _, v := range nums2 {
		nums1 = append(nums1, v)
	}

	sort.Ints(nums1)
}
