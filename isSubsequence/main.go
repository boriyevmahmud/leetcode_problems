package main

import "fmt"

func main() {

	fmt.Println(isSubsequence("abc", "ahbgdc"))
}

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	i, j := 0, 0
	for ; j < len(t); j++ {
		if i < len(s) && t[j] == s[i] {
			i++
		}
	}

	return i == len(s)
}
