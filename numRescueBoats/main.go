package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		people = []int{3, 4, 3}
		limit  = 5
	)
	fmt.Println(numRescueBoats(people, limit))

}

func numRescueBoats(people []int, limit int) int {

	sort.Ints(people)
	counter := 0
	left := 0
	right := len(people) - 1

	for left <= right {

		counter++
		if people[left]+people[right] <= limit {
			left++
		}
		right--
	}

	return counter
}
