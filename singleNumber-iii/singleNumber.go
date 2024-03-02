package main

import "fmt"

func main() {

	fmt.Println(singleNumber([]int{1, 2, 1, 3, 2, 5}))
}

func singleNumber(nums []int) []int {

	response := []int{}
	mapsNum := make(map[int]int)
	for _, num := range nums {
		mapsNum[num] = mapsNum[num] + 1
	}
	for num, val := range mapsNum {
		if val == 1 {
			response = append(response, num)
		}
	}

	return response
}

func singleNumber(nums []int) []int {
	numsMap := map[int]int{}
    result:=[]int{}
	i := 0
	for ; i < len(nums); i++ {
		numsMap[nums[i]]++
	}
	for key, value := range numsMap {
		if value == 1 {
			result=append(result,key)
            
            
        	}
	}
	
	return result
}