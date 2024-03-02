package main

import "fmt"

func main() {

	fmt.Println(singleNumber([]int{2,2,1}))
}

func singleNumber(nums []int) int {

	mapsNum := make(map[int]int)
	for _, num := range nums {
		mapsNum[num] = mapsNum[num] + 1
	}
	for num, val := range mapsNum {
		if val == 1 {
			return num
		}
	}

	return 0
}
