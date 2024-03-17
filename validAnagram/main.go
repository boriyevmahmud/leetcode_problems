package main

import "fmt"

func main() {
	fmt.Println(isAnagram(""))

}
func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	n := len(s) / 2

	for i := 0; i < n; i++ {
		if s[i] != t[n-i] {
			return false
		}
	}

	return true
}
