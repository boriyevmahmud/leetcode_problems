package main

import (
	"fmt"
)

func main() {

	fmt.Println(countBits(5))

}

func countBits(n int) []int {

	resp := []int{}
	for i := 0; i <= n; i++ {
		// formatted := strconv.FormatUint(uint64(i), 2)
		temp := 0
		fmt.Println("i", i)
		fmt.Println(i >>1)
		resp = append(resp, temp)
	}

	return resp
}

// func countBits(n int) []int {

// 	resp := []int{}
// 	for i := 0; i <= n; i++ {
// 		formatted := strconv.FormatUint(uint64(i), 2)
// 		temp := 0
// 		for j := 0; j < len(formatted); j++ {
// 			if string(formatted[j]) == "1" {
// 				temp += 1

// 			}
// 		}
// 		resp = append(resp, temp)
// 	}

// 	return resp
// }
