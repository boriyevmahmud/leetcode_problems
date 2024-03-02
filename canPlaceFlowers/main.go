package main

import "fmt"

func main() {
	fmt.Println(canPlaceFlowers([]int{0, 0, 0, 0}, 3))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i, plot := range flowerbed {
		if plot == 0 {
			continue
		}
		if (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--

		}
	}

	return n <= 0
}
