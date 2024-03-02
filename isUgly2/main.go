package main

import "fmt"

func main() {

	fmt.Println(nthUglyNumber(10))
}

func nthUglyNumber(n int) int {

	counter := 0
	i := 1
	for {
		if ok := isUgly(i); ok {
			counter++
			if counter == n {
				return i
			}
		}
		i++
	}

}

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	for n%2 == 0 {
		n /= 2
	}

	for n%3 == 0 {
		n /= 3
	}

	for n%5 == 0 {
		n /= 5
	}

	return n == 1
}
