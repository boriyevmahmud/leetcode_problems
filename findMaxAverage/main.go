package main

import (
	"fmt"
)

func main() {

	fmt.Println(findMaxAverage([]int{-1}, 1))

}

func findMaxAverage(nums []int, k int) float64 {
	// avgs := []float64{}
	max := 0.0

	for i := 0; i <= len(nums)-k; i++ {
		var sum int
		for j := i; j < k+i; j++ {
			sum += nums[j]
		}
		if max < float64(sum)/float64(k) {
			max = float64(sum) / float64(k)
		} else if float64(sum)/float64(k) < 0 {
			max = float64(sum) / float64(k)
		}
	}

	return max
}
