package main

import "fmt"

func main() {
	fmt.Println(distributeCandies([]int{1, 1, 2, 2, 3, 3}))
}

func distributeCandies(candyType []int) int {
	cnt := 0
	unqType := make(map[int]bool, len(candyType))

	for _, cand := range candyType {
		if !unqType[cand] {
			unqType[cand] = true
			cnt++
		}
	}
	if cnt < len(candyType)/2 {
		return cnt
	} else {
		return len(candyType) / 2
	}

	return 0
}
