package main

import "fmt"

func main() {

	fmt.Println(firstBadVersion(5))
}

func firstBadVersion(n int) int {
	i := 0
	for i <= n {
		if isBadVersion(i) {
			return i
		}
		i += 1
	}
	return 0
}

func isBadVersion(version int) bool {

	return version == 4
}
