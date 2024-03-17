package main

import "fmt"

func main() {
	fmt.Println(findRestaurant(
		[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		[]string{"KFC", "Shogun", "Burger King"},
	))
}

func findRestaurant(list1 []string, list2 []string) []string {
	listsMap := make(map[string]int)
	listsMapExists := make(map[string]int)

	for i, v := range list1 {
		listsMapExists[v] += i
	}
	for i, v := range list2 {
		if val, ok := listsMapExists[v]; ok {
			listsMap[v] = i + val
		}
	}

	min := 100000
	for _, v := range listsMap {
		if v < min {
			min = v
		}
	}
	fmt.Println("min: ", min)
	fmt.Println("min: ", min)

	resp := make([]string, 0, len(listsMap))

	for key, v := range listsMap {
		if v == min {
			resp = append(resp, key)
		}
	}

	return resp
}
