package main

import "fmt"

func main() {
	fmt.Println(isPowerOfTwo(16))
}

func isPowerOfTwo(n int) bool {
	tmp := 1
	for tmp <= n {
		if tmp == n {
			return true
		}
		tmp *= 2
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
