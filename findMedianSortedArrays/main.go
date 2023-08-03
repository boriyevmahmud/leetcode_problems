package main

import "fmt"

func main() {
	a := []int{1, 3}
	b := []int{2}
	fmt.Println(findMedianSortedArrays(a, b))
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mergedArr := append(nums1, nums2...)
	sortedArr := MergeSort(mergedArr)
	resp := 0.0
	if len(sortedArr)%2 == 0 {
		a := len(sortedArr) / 2
		resp = float64(sortedArr[a-1]+sortedArr[a]) / 2
		fmt.Println("sorted array", sortedArr)
		fmt.Println("a", a)
	} else {
		resp = float64(sortedArr[len(sortedArr)/2])
	}

	return resp
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	mergedArr := make([]int, 0)

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(mergedArr, right...)
		} else if len(right) == 0 {
			return append(mergedArr, left...)
		} else if left[0] < right[0] {
			mergedArr = append(mergedArr, left[0])
			left = left[1:]
		} else {
			mergedArr = append(mergedArr, right[0])
			right = right[1:]
		}
	}

	return mergedArr
}

// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	median := 0
// 	for _, elem := range nums1 {
// 		median += elem
// 	}
// 	for _, elem := range nums2 {
// 		median += elem
// 	}

// 	a := float64(median) / float64((len(nums1) + len(nums2)))

// 	return a
// }
