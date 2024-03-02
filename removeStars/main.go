package main

import "fmt"

func main() {

	fmt.Println(removeStars("leet**cod*e"))
}

func removeStars(s string) string {
	resp := []rune{}
	for _, r := range s {
		if r != '*' {
			resp = append(resp, r)
		} else if len(resp) > 0 {
			resp = resp[:len(resp)-1]
		}
	}

	return string(resp)
}
