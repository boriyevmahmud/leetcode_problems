package main

import "fmt"

func main() {
	var numb int
	fmt.Scan(&numb)
	fmt.Println(isPalindrome(numb))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	num := x

	remainder := 0
	res := 0
	for x > 0 {
		remainder = x % 10
		res = (res * 10) + remainder
		x = x / 10
	}

	return res == num
}
