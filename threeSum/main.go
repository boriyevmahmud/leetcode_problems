package main

import "fmt"

func main() {

	fmt.Println(threeSum([]int{0, 3, 0, 1, 1, -1, -5, -5, 3, -3, -3, 0}))
}

func threeSum(nums []int) [][]int {

	resp := [][]int{}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					arr := []int{nums[i], nums[j], nums[k]}
					fmt.Println(arr)
					if !check(arr, resp) {
						resp = append(resp, arr)
					}
					fmt.Println("resp: ", resp)
				}
			}
		}
	}

	return resp
}

func check(arr []int, bigArr [][]int) bool {
	arrMap := make(map[int]bool)
	res := false
	for _, i := range arr {
		arrMap[i] = true
	}
	fmt.Println("arrMap: ", arrMap)
	for i := 0; i < len(bigArr); i++ {
		if arrMap[bigArr[i][0]] && arrMap[bigArr[i][1]] && arrMap[bigArr[i][2]] {
			res = true
			return res
		}
	}
	fmt.Println("in here",res)

	return res
}
