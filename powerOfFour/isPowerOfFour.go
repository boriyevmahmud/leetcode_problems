package main

import (
	"fmt"
)

func main() {

	fmt.Println(isPowerOfFour(5))
}

func isPowerOfFour(n int) bool {
	if n == 1 {
		return true
	}
	tmpN := float64(n)
	for tmpN > 1 {
		tmpN = tmpN / 4
		if tmpN == 1 {
			return true
		}
	}

	return false
}
