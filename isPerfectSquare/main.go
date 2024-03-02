package main

import "fmt"

func main() {

	fmt.Println(isPerfectSquare(14))

}

func isPerfectSquare(num int) bool {
	for i := 1; i <= num; i++ {
		if i*i == num {
			return true
		}
	}

	return false
}
