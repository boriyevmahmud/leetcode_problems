package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {
	if x <0 {
	return false
	}
	reversed := 0
	x1 := x
	for x > 0 {
		reversed = reversed*10 + x % 10
		x = x / 10
	}

	return x1 == reversed
}
