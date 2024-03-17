package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(integerBreak(10))
}

func integerBreak(n int) int {

	if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	}

	countOf3s := n / 3
	remainder := n % 3
	if remainder == 0 {
		return int(math.Pow(3, float64(countOf3s)))
	} else if remainder == 1 {
		return int(math.Pow(3, float64(countOf3s-1))) * 4
	} else {
		return int(math.Pow(3, float64(countOf3s))) * 2
	}

}
