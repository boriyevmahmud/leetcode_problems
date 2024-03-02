package main

import "fmt"

func main() {

	fmt.Println(addDigits(38))
}

// func addDigits(num int) int {

// 	if num >= 0 && num < 10 {
// 		return num
// 	}
// 	resp := 0
// 	var tmp = num
// 	tmp /= 10
// 	a := 0
// 	for tmp != 0 {
// 		a = tmp % 10
// 		fmt.Println("a", a)
// 		resp += a
// 		tmp /= 10
// 	}

// 	return 1
// }

func addDigits(num int) int {
	if num < 10 {
		return num
	}
	if num%9 == 0 {
		return 9
	}
	return num % 9
}
