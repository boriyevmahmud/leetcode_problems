package main

import "fmt"

func main() {
	fmt.Println(isPowerOfThree(2))
}

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	tmp := 1
	for tmp <= n {
		if tmp == n {
			return true
		}
		tmp *= 3
	}

	return false
}

// func isPowerOfTwo(n int) bool {
// 	floatN := float32(n)
// 	for floatN != 0 {
// 		a := n % 2
// 		floatN /= 2
// 		if a == 0 {
// 			return true
// 		}
// 	}

// 	return false
// }
