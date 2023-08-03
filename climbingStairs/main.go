package main

import "fmt"

func main() {
	fmt.Println(climbStairs(4))
}

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}

	arr := make([]int, n+1)
	arr[1] = 1
	arr[2] = 2

	for i := 3; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	return arr[n]
}
