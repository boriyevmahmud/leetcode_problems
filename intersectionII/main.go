package main

import "fmt"

func main() {

	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
}

func intersect(nums1 []int, nums2 []int) []int {
	numsMap := make(map[int]int)
	result := make([]int, 0, len(nums1)+len(nums2))

	for _, n := range nums1 {
		numsMap[n]++
	}

	for _, n := range nums2 {
		if numsMap[n] > 0 {
			result = append(result, n)
			numsMap[n]--
		}

	}

	return result
}

// func intersect(nums1, nums2 []int) []int {

// 	resp := []int{}
// 	map1 := make(map[int]int)
// 	map2 := make(map[int]int)
// 	for _, num := range nums1 {
// 		map1[num] += 1
// 	}

// 	for _, num := range nums2 {
// 		map2[num] += 1
// 	}

// 	for k, v := range map1 {
// 		for k1, v1 := range map2 {
// 			if k1 == k {
// 				var min = 0
// 				if v1 == v || v1 >= v {
// 					min = v
// 				} else if v1 <= v {
// 					min = v1
// 				}
// 				for min != 0 {
// 					resp = append(resp, k)
// 					min -= 1
// 				}
// 			}
// 		}

// 	}
// 	return resp
// }
