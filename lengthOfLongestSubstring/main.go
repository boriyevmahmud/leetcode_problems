package main

import (
	"fmt"
)

func main() {
	var (
		people = "dvdf"
	)
	fmt.Println(lengthOfLongestSubstring(people))

}

func lengthOfLongestSubstring(s string) int {

	if len(s) == 1 {
		return 1
	}
	max := 0

	for i := 0; i < len(s); i++ {
		check := make(map[byte]bool)
		check[s[i]] = true
		for j := i + 1; j < len(s); j++ {

			if check[s[j]] {
				break
			} else {
				check[s[j]] = true
			}
		}

		max = findMax(max, len(check))
	}

	return max
}

func findMax(nums ...int) int {
	maxNum := nums[0]

	for _, num := range nums[1:] {
		if maxNum < num {
			maxNum = num
		}
	}

	return maxNum
}
