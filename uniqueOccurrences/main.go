package main

import "fmt"

func main() {

	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
}
func uniqueOccurrences(arr []int) bool {
	frequencyMap := make(map[int]int)
	occurrencesMap := make(map[int]bool)

	// Count the occurrences of each value
	for _, v := range arr {
		frequencyMap[v]++
	}

	// Check if the occurrences are unique
	for _, count := range frequencyMap {
		if occurrencesMap[count] {
			return false
		}
		occurrencesMap[count] = true
	}

	return true
}
