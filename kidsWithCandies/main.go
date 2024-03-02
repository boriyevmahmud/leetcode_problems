package main

import "fmt"

func main() {

	fmt.Println(kidsWithCandies([]int{12, 1, 12}, 10))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {

	resp := make([]bool, 0, len(candies))
	maxCandies := 0

	for _, v := range candies {
		if v > maxCandies {
			maxCandies = v
		}
	}
	for _, v := range candies {
		resp = append(resp, v+extraCandies >= maxCandies)
	}

	return resp
}
