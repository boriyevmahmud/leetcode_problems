package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}

func topKFrequent(nums []int, k int) []int {
	type St struct {
		number int
		count  int
	}
	elements := []St{}
	numsMap := make(map[int]int)

	for _, num := range nums {
		numsMap[num] = numsMap[num] + 1
	}
	for k, v := range numsMap {
		elements = append(elements, St{
			number: k,
			count:  v,
		})
	}
	sort.Slice(elements, func(i, j int) bool {
		if elements[i].count < elements[j].count {
			return true
		} else if elements[i].count == elements[j].count {
			return elements[i].number < elements[j].number
		} else {
			return false
		}
	})
	res := make([]int, 0, k)
	for i := 0; i < k && i < len(elements); i++ {
		res = append(res, elements[i].number)
	}
	fmt.Println("element", elements)

	return res
}
